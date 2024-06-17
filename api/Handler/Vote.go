package Handler

import (
	"auth-api/Database"
	"auth-api/Form"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListVote(c *fiber.Ctx) error {
	var listbody Form.ListVote

	if err := c.BodyParser(&listbody); err != nil {
		return err
	}

	if listbody.Token != "" {
		data ,err := Database.ListVote(listbody.Token)
		fmt.Println(data, err)
		if err != nil {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})

		} else {	
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Listed for vote successfully",
				"data": data,
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": 502,
			"error":  "İnput is incorrect",
		})
	}
	return nil
}


func VoteAdd(c *fiber.Ctx) error {
	var addbody Form.AddVote

	if err := c.BodyParser(&addbody); err != nil {
		return err
	}

	if addbody.Voter != "" || addbody.Score != 0 || addbody.Anime != "" {
		err := Database.AddVote(addbody.Anime, addbody.Voter, addbody.Score)
		if !err {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Added for vote successfully",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": 502,
			"error":  "İnput is incorrect",
		})
	}
	return nil
}


func VoteDel(c *fiber.Ctx) error {
	var delVote Form.DelVote

	if err := c.BodyParser(&delVote); err != nil {
		return err
	}

	if delVote.Voter != "" || delVote.Anime != "" {
		err := Database.DelVote(delVote.Voter, delVote.Anime, delVote.Score)
		if !err {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Delete for anime successfully",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": 502,
			"error":  "İnput is incorrect",
		})
	}
	return nil
}