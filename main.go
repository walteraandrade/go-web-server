package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	"go-web-server/models"
	"go-web-server/db"
	"go-web-server/config"
	"database/sql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func createAlbumHandler(w http.ResponseWriter, r *http.Request, albumDb *sql.DB) {
        var album models.Album

        err := json.NewDecoder(r.Body).Decode(&album)
        if err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return
        }

        fmt.Printf("Received album: %+v\n", album)
        err = db.InsertAlbum(albumDb, album)
        if err != nil {
            http.Error(w, "Error inserting album", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, "Album created successfully")
}

func main() {
	env_manager.LoadEnv()

	http.HandleFunc("/", handler)
	http.HandleFunc("/insert-album", func(w http.ResponseWriter, r *http.Request) {
	albumDb, err := db.DbConnect()
	if err != nil {
        fmt.Println("Error connecting to the database: ", err)
    }
    defer albumDb.Close()
		createAlbumHandler(w, r, albumDb)
	})
	http.ListenAndServe(":8080", nil)
}
