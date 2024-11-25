package app

import (
	"fmt"
	"os"

	"github.com/avehun/songlib/internal/pkg/handlers"
	"github.com/avehun/songlib/internal/pkg/repo"
	"github.com/avehun/songlib/internal/pkg/routers"
	"github.com/avehun/songlib/internal/pkg/servers"
	"github.com/avehun/songlib/internal/pkg/services"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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
	DATABASE_URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// m, err := migrate.New(
	// 	"/home/avehun/go/github.com/avehun/songlib/migrations",
	// 	DATABASE_URL,
	// )
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("Ошибка применения миграции: %v", err)
	// }

	// if err != nil {
	// 	log.Errorf("Error creating migrations: %v", err)
	// }
	db, err := sqlx.Open("pgx", DATABASE_URL)
	if err != nil {
		log.Errorf("Error connetcing to data base: %v", err)
	}
	defer db.Close()

	repo := repo.NewSongRepo(db)
	service := services.NewSongService(repo)
	handler := handlers.NewHttpHandler(service)
	r := routers.NewRouter(*handler)
	serv := servers.NewHttpServer(r)
	serv.Start()
}
