package Handler

import (
	"auth-api/Database"
	"auth-api/Form"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AnimeAdd(c *fiber.Ctx) error {
	var addbody Form.AddAnime

	if err := c.BodyParser(&addbody); err != nil {
		return err
	}
	fmt.Println(addbody)

	if addbody.Name != "" || addbody.Description != "" || addbody.Thumbnail != "" {
		err := Database.InsertAnime(addbody.Name, addbody.Thumbnail, addbody.Description)
		if !err {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Added for anime successfully",
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

func AnimeDel(c *fiber.Ctx) error {
	var delAnime Form.DelAnime

	if err := c.BodyParser(&delAnime); err != nil {
		return err
	}

	if delAnime.Name != "" {
		err := Database.DeleteAnime(delAnime.Name)
		if err != nil {
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

func AnimeList(c *fiber.Ctx) error {
	data, err := Database.ListAnime()
	if err != nil {
		return c.JSON(fiber.Map{
			"data":   err,
			"status": "error",
		})
	}
	return c.JSON(fiber.Map{
		"data":   data,
		"status": "OK",
	})
}

func Anime(c *fiber.Ctx) error {
	var Anime Form.Anime

	if err := c.BodyParser(&Anime); err != nil {
		return err
	}

	if Anime.Name != "" {
		data, err := Database.InfoAnime(Anime.Name)
		if err != nil {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Info for anime successfully",
				"data":    data,
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

func Search(c *fiber.Ctx) error {
	var Search Form.Search

	if err := c.BodyParser(&Search); err != nil {
		return err
	}

	data, _ := Database.Search(Search.Query)

	c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Info for anime successfully",
		"data":    data,
	})

	return nil
}
