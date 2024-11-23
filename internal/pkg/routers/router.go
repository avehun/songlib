package routers

import (
	"net/http"

	"github.com/avehun/songlib/internal/pkg/handlers"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.ListRoutes)
	r.HandleFunc("GET /songs/", handlers.ListSongs)
	r.HandleFunc("GET /songs/{id}/", handlers.RetrieveSong)
	r.HandleFunc("DELETE /songs/{id}/", handlers.DeleteSong)
	r.HandleFunc("PUT /songs/{id}/", handlers.ChangeSong)
	r.HandleFunc("POST /songs/", handlers.AddSong)

	return r
}
