package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS movies (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	year INTEGER,
	rating REAL
);

CREATE TABLE IF NOT EXISTS genres (
	id INTEGER PRIMARY KEY,
	movie_id INTEGER,
	genre TEXT,
	FOREIGN KEY (movie_id) REFERENCES movies (id)
);
`

func main() {
	// Open database
	db, err := sql.Open("sqlite", "./movies.db")
	if err != nil {
		log.Fatalf("Sry,There is an error opening database: %v", err)
	}
	defer db.Close()

	// Create tables
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Sry,There is an error creating schema: %v", err)
	}

	fmt.Println("Database schema created successfully.")

	populateDatabase(db, "./IMDB-movies.csv", "./IMDB-movies_genres.csv")

	exampleQuery(db)

	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func populateDatabase(db *sql.DB, moviesFile string, genresFile string) {

	_, err := db.Exec("DELETE FROM movies; DELETE FROM genres;")
	if err != nil {
		log.Fatalf("Sry,There is an error clearing tables: %v", err)
	}
	fmt.Println("Existing data cleared.")

	moviesCSV, err := os.Open(moviesFile)
	if err != nil {
		log.Fatalf("Sry,There is an error opening movies file: %v", err)
	}
	defer moviesCSV.Close()

	moviesReader := csv.NewReader(bufio.NewReader(moviesCSV))
	moviesReader.Read()

	for {
		record, err := moviesReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Sry,There is an error reading movies CSV: %v", err)
		}

		id, _ := strconv.Atoi(record[0])
		title := record[1]
		year, _ := strconv.Atoi(record[2])
		rating, _ := strconv.ParseFloat(record[3], 64)

		var exists int
		err = db.QueryRow("SELECT COUNT(*) FROM movies WHERE id = ?", id).Scan(&exists)
		if err != nil {
			log.Fatalf("Sry,There is an error checking existing record: %v", err)
		}

		if exists > 0 {
			fmt.Printf("Skipping duplicate id: %d\n", id)
			continue
		}

		_, err = db.Exec("INSERT INTO movies (id, title, year, rating) VALUES (?, ?, ?, ?)", id, title, year, rating)
		if err != nil {
			log.Fatalf("Sry,There is an error inserting into movies table: %v", err)
		}
	}

	genresCSV, err := os.Open(genresFile)
	if err != nil {
		log.Fatalf("Sry,There is an error opening genres file: %v", err)
	}
	defer genresCSV.Close()

	genresReader := csv.NewReader(bufio.NewReader(genresCSV))
	genresReader.Read()

	for {
		record, err := genresReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Sry,There is an error reading genres CSV: %v", err)
		}

		id, _ := strconv.Atoi(record[0])
		genre := record[1]

		_, err = db.Exec("INSERT INTO genres (movie_id, genre) VALUES (?, ?)", id, genre)
		if err != nil {
			log.Fatalf("Sry,There is an error inserting into genres table: %v", err)
		}
	}

	fmt.Println("Data populated successfully.")
}

func exampleQuery(db *sql.DB) {
	query := `
	SELECT movies.title, genres.genre, movies.rating
	FROM movies
	JOIN genres ON movies.id = genres.movie_id
	ORDER BY movies.rating DESC
	LIMIT 5
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Sry,There is an error running query: %v", err)
	}
	defer rows.Close()

	fmt.Println("Top Movies by Rating and Genre:")
	for rows.Next() {
		var title, genre string
		var rating float64
		if err := rows.Scan(&title, &genre, &rating); err != nil {
			log.Fatalf("Sry,There is an error scanning row: %v", err)
		}
		fmt.Printf("Title: %s, Genre: %s, Rating: %.1f\n", title, genre, rating)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Sry,There is an error with rows: %v", err)
	}
}
