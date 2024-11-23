package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/avehun/songlib/internal/pkg/services"
)

func ListRoutes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GET /songs/\n")
	io.WriteString(w, "GET /songs/{id}\n")
	io.WriteString(w, "DELETE /songs/{id}\n")
	io.WriteString(w, "PUT /songs/{id}\n")
	io.WriteString(w, "POST /songs/\n")
}
func ListSongs(w http.ResponseWriter, r *http.Request) {
	songs := services.ListSongs()
	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, "Server error", 400)
	}
}
func RetrieveSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	var song models.Song
	song = services.RetrieveSong(id)
	json.NewEncoder(w).Encode(song)
}
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		services.DeleteSong(id)
	}
}
func ChangeSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		services.ChangeSong(id)
	}
}
func AddSong(w http.ResponseWriter, r *http.Request) {
	song := models.Song{}
	req, err := http.NewRequest("GET", os.Getenv("OUTER_SERVICE_URL"), r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	json.NewDecoder(res.Body).Decode(&song)

	services.AddSong(song)
}

func parseId(url string) (string, error) {
	splitPath := strings.Split(url, "/")
	if len(splitPath) < 3 {
		return "", errors.New("Id required")
	}
	return splitPath[2], nil
}
