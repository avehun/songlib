package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/avehun/songlib/internal/pkg/services"
	log "github.com/sirupsen/logrus"
)

func ListRoutes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GET /songs/\n")
	io.WriteString(w, "GET /songs/{id}\n")
	io.WriteString(w, "DELETE /songs/{id}\n")
	io.WriteString(w, "PUT /songs/{id}\n")
	io.WriteString(w, "POST /songs/\n")
}

// Songlib godoc
// @Summary      Get songs
// @Description  List all songs with pagination and filtering
// @Tags         songs
// @Produce      json
// @Success      200  {object}  []models.Song
// @Router       /songs/ [get]
func ListSongs(w http.ResponseWriter, r *http.Request) {
	songs := services.ListSongs()
	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, "Server error", 400)
		return
	}
}

// Songlib godoc
// @Summary      Retrieve song
// @Description  Retrieve a song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Song ID"
// @Success      200  {object}  models.Song
// @Router       /songs/{id} [get]
func RetrieveSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	song := services.RetrieveSong(id)
	json.NewEncoder(w).Encode(song)
}

// Songlib godoc
// @Summary      Delete song
// @Description  Delete a song by id
// @Tags         songs
// @Param        id   path      int  true  "Song ID"
// @Success      200
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	services.DeleteSong(id)
}

// Songlib godoc
// @Summary      Change song
// @Description  Change a song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Song ID"
// @Param        song body models.Song true "change existing song"
// @Success      200
// @Router       /songs/{id} [patch]
func ChangeSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	services.ChangeSong(id)

}

// Songlib godoc
// @Summary      Add song
// @Description  Add a new song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        song body Song true "add new song"
// @Success      200
// @Router       /songs/ [post]
func AddSong(w http.ResponseWriter, r *http.Request) {
	song := models.Song{}
	json.NewDecoder(r.Body).Decode(&song)

	params := url.Values{}
	params.Add("group", song.Group)
	params.Add("song", song.Song)

	uri := fmt.Sprintf("%s/info?%s", os.Getenv("OUTER_SERVICE_URL"), params.Encode())

	resp, err := http.Get(uri)

	if err != nil {
		log.Info("Unable to fetch from outer service")
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&song)

	services.AddSong(song)
}

func parseId(url string) (string, error) {
	splitPath := strings.Split(url, "/")
	if len(splitPath) < 3 {
		return "", errors.New("id required")
	}
	return splitPath[2], nil
}

type Song struct {
	Group string `json="group`
	Song  string `json="song"`
}
