package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/avehun/songlib/internal/pkg/services"
	log "github.com/sirupsen/logrus"
)

type HttpHandler struct {
	songService *services.SongService
}

func NewHttpHandler(service *services.SongService) *HttpHandler {
	return &HttpHandler{
		songService: service,
	}
}

// Songlib godoc
// @Summary      Get songs
// @Description  List all songs with pagination and filtering
// @Tags         songs
// @Produce      json
// @Success      200  {object}  []models.Song
// @Router       /songs/ [get]
func (h *HttpHandler) ListSongs(w http.ResponseWriter, r *http.Request) {
	songs := h.songService.ListSongs()
	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, "Server error", 400)
		return
	}
	w.WriteHeader(http.StatusOK)
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
func (h *HttpHandler) RetrieveSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	song := h.songService.RetrieveSong(id)
	json.NewEncoder(w).Encode(song)
	w.WriteHeader(http.StatusOK)
}

// Songlib godoc
// @Summary      Delete song
// @Description  Delete a song by id
// @Tags         songs
// @Param        id   path      int  true  "Song ID"
// @Success      200
// @Router       /songs/{id} [delete]
func (h *HttpHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	h.songService.DeleteSong(id)
	w.WriteHeader(http.StatusOK)
}

// Songlib godoc
// @Summary      Change song
// @Description  Change a song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        song body models.Song true "change existing song"
// @Success      200
// @Router       /songs/{id} [patch]
func (h *HttpHandler) ChangeSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	h.songService.ChangeSong(song)
	w.WriteHeader(http.StatusOK)
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
func (h *HttpHandler) AddSong(w http.ResponseWriter, r *http.Request) {
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

	h.songService.AddSong(song)
	w.WriteHeader(http.StatusOK)
}

func parseId(url string) (string, error) {
	splitPath := strings.Split(url, "/")
	if len(splitPath) < 3 {
		return "", errors.New("id required")
	}
	return splitPath[2], nil
}

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
