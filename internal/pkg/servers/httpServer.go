package servers

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type httpServer struct {
	router *http.ServeMux
}

func NewHttpServer(router *http.ServeMux) *httpServer {
	return &httpServer{
		router: router,
	}
}

func (serv *httpServer) Start() {
	log.Infof("HTTP servser listening on port %v", "localhost:"+os.Getenv("HTTP_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), serv.router)
	if err != nil {
		log.Fatalf("Starting server error occured: %v", err)
	}
}
