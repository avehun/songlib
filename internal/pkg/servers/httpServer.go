package servers

import (
	"fmt"
	"net/http"
	"os"
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
	err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), serv.router)
	if err != nil {
		fmt.Printf("Starting server error occured: %v", err)
	}
}
