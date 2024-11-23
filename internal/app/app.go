package app

import (
	"github.com/avehun/songlib/internal/pkg/routers"
	"github.com/avehun/songlib/internal/pkg/servers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	log.Info("Initializing application")
	_ = godotenv.Load()
	r := routers.NewRouter()
	serv := servers.NewHttpServer(r)
	serv.Start()
}
