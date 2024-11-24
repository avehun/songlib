package main

import (
	"github.com/avehun/songlib/internal/app"
)

// @title           Songlib documentation
// @version         1.0
// @description     This is an API songs store server.
// @host      localhost:8080
// @BasePath  /
func main() {
	app := app.NewApp()
	app.Run()
}
