package Database

import (
	"auth-api/Form"
	"auth-api/Utils"
	"database/sql"
	"fmt"
)

func InsertAnime(Name, Thumbnail, Description string) bool {
	fmt.Print(Description)
	Token := Utils.Token(5)
	if !FinderAnime(Name) {
		query := "INSERT INTO anime (name, thumbnail, description, token) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(query, Name, Thumbnail, Description, Token)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func FinderAnime(Name string) bool {
	query := "SELECT COUNT(*) FROM anime WHERE name = ?"
	row := db.QueryRow(query, Name)

	var anime int
	err := row.Scan(&anime)
	if err != nil {
		return false
	}

	if anime == 0 {
		return false
	}

	return true
}

func FindAnime(Name string) bool {
	query := "SELECT COUNT(*) FROM anime WHERE token = ?"
	row := db.QueryRow(query, Name)

	var anime int
	err := row.Scan(&anime)
	if err != nil {
		return false
	}

	if anime == 0 {
		return false
	}

	return true
}

func DeleteAnime(name string) error {
	query := "DELETE FROM anime WHERE name = ?"
	result, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Anime Bulunamadı")
	}

	return nil
}

func InfoAnime(name string) (*Form.Anime, error) {
	query := "SELECT token, name, description, thumbnail, five, four, three, two, one FROM anime WHERE name = ?"
	row := db.QueryRow(query, name)
	var anime Form.Anime
	if err := row.Scan(&anime.Token, &anime.Name, &anime.Description, &anime.Thumbnail, &anime.Five, &anime.Four, &anime.Three, &anime.Two, &anime.One); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	anime.Average = Utils.CalculateRating(anime.Five, anime.Four, anime.Three, anime.Two, anime.One)

	return &anime, nil
}

func ListAnime() ([]Form.Anime, error) {
	query :="SELECT token, name, description, thumbnail, five, four, three, two, one FROM anime"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animes []Form.Anime
	for rows.Next() {
		var anime Form.Anime
		if err := rows.Scan(&anime.Token, &anime.Name, &anime.Description, &anime.Thumbnail, &anime.Five, &anime.Four, &anime.Three, &anime.Two, &anime.One); err != nil {
			return nil, err
		}
		anime.Average = Utils.CalculateRating(anime.Five, anime.Four, anime.Three, anime.Two, anime.One)
		animes = append(animes, anime)
	}

	return animes, nil
}



func Search(data string) ([]Form.Anime, error) {
	query := "SELECT name, token, description, thumbnail, five, four, three, two, one FROM anime WHERE name LIKE ?"
	rows, err := db.Query(query, "%"+data+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var animes []Form.Anime
	for rows.Next() {
		var anime Form.Anime
		if err := rows.Scan(&anime.Name, &anime.Token, &anime.Description, &anime.Thumbnail, &anime.Five, &anime.Four, &anime.Three, &anime.Two, &anime.One); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		animes = append(animes, anime)
	}

	return animes, nil
}
