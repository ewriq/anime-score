package Database

import (
	"auth-api/Form"
	"auth-api/Utils"
	"fmt"
)

func AddVote(anime, voter string, score int) bool {
	Token := Utils.Token(10)
	if !FindVoter(voter, anime) {
		query := "INSERT INTO vote (anime, score, voter, token) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(query, anime, score, voter, Token)
		if err != nil {
			fmt.Println(err)
			return false
		}

		var updateAnime string
		switch score {
		case 1:
			updateAnime = "UPDATE anime SET one = one + 1 WHERE name = ?"
		case 2:
			updateAnime = "UPDATE anime SET two = two + 1 WHERE name = ?"
		case 3:
			updateAnime = "UPDATE anime SET three = three + 1 WHERE name = ?"
		case 4:
			updateAnime = "UPDATE anime SET four = four + 1 WHERE name = ?"
		case 5:
			updateAnime = "UPDATE anime SET five = five + 1 WHERE name = ?"
		default:
			fmt.Println("Indivail score:", score)
			return false
		}

		_, err = db.Exec(updateAnime, anime)
		if err != nil {
			fmt.Println("Err Anime table update:", err)
			return false
		}

		return true
	} else {
		return false
	}
}

func FindVoter(voter, anime string) bool {
	query := "SELECT COUNT(*) FROM vote WHERE voter = ? AND anime = ?"
	row := db.QueryRow(query, voter, anime)

	var animes int
	err := row.Scan(&animes)
	if err != nil {
		return false
	}

	if animes == 0 {
		return false
	}

	return true
}

func DelVote(voter, anime string, score int) bool {
	query := "DELETE FROM vote WHERE voter = ? AND anime = ?"
	result, err := db.Exec(query, voter, anime)
	if err != nil {
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	if rowsAffected == 0 {
		return false
	}

	var updateAnime string
		switch score {
		case 1:
			updateAnime = "UPDATE anime SET one = one - 1 WHERE name = ?"
		case 2:
			updateAnime = "UPDATE anime SET two = two - 1 WHERE name = ?"
		case 3:
			updateAnime = "UPDATE anime SET three = three - 1 WHERE name = ?"
		case 4:
			updateAnime = "UPDATE anime SET four = four - 1 WHERE name = ?"
		case 5:
			updateAnime = "UPDATE anime SET five = five - 1 WHERE name = ?"
		default:
			fmt.Println("Indivail score:", score)
			return false
	}

	_, err = db.Exec(updateAnime, anime)
	if err != nil {
		return false
	}

	return true
}

func ListVote(voter string) ([]Form.Vote, error) {
	query := "SELECT token, voter, score, anime FROM vote WHERE voter = ?"
	rows, err := db.Query(query, voter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var votes []Form.Vote
	for rows.Next() {
		var vote Form.Vote
		err := rows.Scan(&vote.Token, &vote.Voter, &vote.Score, &vote.Anime)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(votes) == 0 {
		return nil, fmt.Errorf("hiç oy bulunamadı")
	}

	return votes, nil
}
