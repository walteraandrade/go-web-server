package db

import (
    "database/sql"
    "errors"
    "go-web-server/models"
)

func InsertAlbum(db *sql.DB, album models.Album) error {
    // Assuming a table named "albums" with columns matching Album struct fields
    sqlStatement := `
        INSERT INTO albums (title, release_date, label, artists, producer, genres, image, rating, review, tracks)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

    // Handle tracks separately (or adjust the SQL statement accordingly)
    // ...

    _, err := db.Exec(sqlStatement, album.Title, album.ReleaseDate, album.Label, album.Artists, album.Producer, album.Genres, album.Image, album.Rating, album.Review)
    if err != nil {
            return errors.New("error inserting album: " + err.Error())
    }

    return nil
}
