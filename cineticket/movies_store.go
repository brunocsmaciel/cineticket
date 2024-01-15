package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/brunocsmaciel/cineticket/types"
)

func GetMovies() []*types.Movie {

	db, err := brenis()
	if err != nil {
		log.Fatal("error while connecting to database")
	}

	results, err := db.Query("SELECT * FROM movies")
	if err != nil {
		panic(err)
	}

	var movies []*types.Movie
	for results.Next() {
		var movie types.Movie
		var actorsJSON string
		var languagesJSON string
		var genreJSON string
		var premiereDateStr string

		err = results.Scan(
			&movie.ID,
			&movie.Name,
			&movie.RunningTime,
			&actorsJSON,
			&languagesJSON,
			&genreJSON,
			&premiereDateStr)

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(actorsJSON), &movie.Actors)
		if err != nil {
			fmt.Println("Error converting actors JSON:", err)
			return nil
		}

		err = json.Unmarshal([]byte(languagesJSON), &movie.Languages)
		if err != nil {
			fmt.Println("Error converting actors JSON:", err)
			return nil
		}

		err = json.Unmarshal([]byte(genreJSON), &movie.Genre)
		if err != nil {
			fmt.Println("Error converting actors JSON:", err)
			return nil
		}

		layoutStrings := []string{"2020-01-14", "2020-14-01 23:21:05"}

		for _, layout := range layoutStrings {
			movie.PremiereDate, err = time.Parse(layout, premiereDateStr)
			if err == nil {
				break // Successfully parsed using the current layout
			}
		}

		movies = append(movies, &movie)
	}

	return movies
}
