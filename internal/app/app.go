package app

import (
	"github.com/avehun/songlib/internal/pkg/routers"
	"github.com/avehun/songlib/internal/pkg/servers"
	"github.com/joho/godotenv"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	_ = godotenv.Load()
	r := routers.NewRouter()
	serv := servers.NewHttpServer(r)
	serv.Start()
}
