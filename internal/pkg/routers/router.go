package routers

import (
	"net/http"
	"os"

	_ "github.com/avehun/songlib/docs"
	"github.com/avehun/songlib/internal/pkg/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", httpSwagger.Handler(httpSwagger.URL("http://localhost:"+os.Getenv("HTTP_PORT")+"/swagger/doc.json")))
	r.HandleFunc("GET /songs/", handlers.ListSongs)
	r.HandleFunc("GET /songs/{id}/", handlers.RetrieveSong)
	r.HandleFunc("DELETE /songs/{id}/", handlers.DeleteSong)
	r.HandleFunc("PUT /songs/{id}/", handlers.ChangeSong)
	r.HandleFunc("POST /songs/", handlers.AddSong)

	return r
}
