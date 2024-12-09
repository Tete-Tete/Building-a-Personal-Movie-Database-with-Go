# Building a Personal Movie Database with Go

This project involved creating a local SQLite database to manage movie data using Go. The database schema was defined to accommodate two primary tables:
- **movies Table**: Stores information about movies, including their unique ID, title, release year, and rating.

- **genres Table**: Stores genres associated with movies, linking each genre to a movie via a foreign key.

The database schema was implemented programmatically using Go and was populated with data from two CSV files:
- **IMDB-movies.csv**::Contains movie details.

- **IMDB-movies_genres.csv** :Contains movie genres.

## Requirements
- Go programming language installed (version 1.16 or higher recommended).

## Installation From Git and Set up for your own computer
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/Building-a-Personal-Movie-Database-with-Go.git
```

### Step 2: Run the Application
To run the Go application, run the following command in your terminal:
```sh
go build -o database.exe main.go
```
This will create a file named `database` in your current directory.

## Key Step
1. SQLite database setup was performed using the `modernc.org/sqlite driver`, which ensures compatibility with Go.
2. Tables were created with appropriate constraints to enforce data integrity.
3. Data was inserted into the tables after ensuring no duplication using `SELECT COUNT(*)` checks.
4. A query was implemented to join the `movies` and `genres` tables, showcasing the top-rated movies and their associated genres.
The resulting database enables efficient querying of movie information, laying the groundwork for building a useful movie management application.


## Purpose and User Interactions
### Purpose:
The personal movie database serves as a tailored solution for movie enthusiasts to organize and rate their collections. It provides capabilities beyond IMDb by integrating personal data and preferences.
### User Interactions:
- **Adding Movies to Personal Collection**: Users can add movies from the existing `movies` table to their personal collection, specifying storage locations and personal ratings.
- **Querying Personal Data**: Users can search their collection based on various filters.
- **Custom Reports**: Generate reports such as "Top 10 movies I own" or "Movies to watch next."
The application could also feature a simple command-line interface (CLI) or a web-based UI for enhanced user experience

## Advantages Over IMDb
Unlike IMDb, which is primarily a read-only database of general movie information, this application:
1. Supports Personalization: Tracks user-specific data, such as storage locations and personal ratings.
2. Offline Access: Operates locally without requiring internet connectivity.
3. Custom Queries: Allows users to generate tailored queries that reflect their personal movie preferences.
4. Integrated Collection Management: Combines movie metadata with personal organization needs.

## Possible Database Enhancements
### Movie Reviews Table:
A table to store personal reviews:
```sh
CREATE TABLE reviews (
    id INTEGER PRIMARY KEY,
    movie_id INTEGER NOT NULL,
    review TEXT,
    FOREIGN KEY (movie_id) REFERENCES movies (id)
);
```
This would enable users to maintain a detailed log of their thoughts on each movie.
### Watchlist Integration:
Add a `watchlist` table to track movies users plan to watch.

## Further Application Development:
1. Recommendation System: Based on personal ratings and reviews, suggest movies from the database or external sources.
2. Mobile App Integration: Develop a mobile version of the application to access and update the collection on the go.
3. Data Import/Export: Support importing/exporting data to/from other movie collection tools.
4. Analytics Dashboard: Provide insights like "Average ratings by genre" or "Most popular directors in my collection."

By incorporating these enhancements, the application can evolve into a comprehensive movie management and analytics platform tailored to individual users.

