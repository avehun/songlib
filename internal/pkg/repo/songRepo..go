package repo

import (
	"fmt"

	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type SongRepo struct {
	db *sqlx.DB
}

// Конструктор для SongRepo
func NewSongRepo(db *sqlx.DB) *SongRepo {

	song := SongRepo{db: db}
	err := song.db.Ping()
	if err != nil {
		log.Fatalf("Cant reach data base: %v", err)
	}
	return &song

}

func (r *SongRepo) GetById(id string) (*models.Song, error) {
	query := `SELECT id, "group", song, release_date, "text", link FROM "songs" WHERE id = $1`
	var song models.Song
	err := r.db.Get(&song, query, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching song by id: %w", err)
	}
	return &song, nil
}

func (r *SongRepo) GetAll() ([]models.Song, error) {
	query := `SELECT id, "group", song, release_date, text, link FROM "songs"`
	var songs []models.Song
	err := r.db.Select(&songs, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all songs: %w", err)
	}
	return songs, nil
}

func (r *SongRepo) Create(song *models.Song) error {
	query := `
		INSERT INTO "songs" ("group", song, release_date, text, link)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	err := r.db.QueryRow(
		query,
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
	).Scan(&song.ID)
	if err != nil {
		return fmt.Errorf("error creating song: %w", err)
	}
	return nil
}

func (r *SongRepo) Update(song *models.Song) error {
	query := `
		UPDATE "songs"
		SET "group" = $1, song = $2, release_date = $3, text = $4, link = $5
		WHERE id = $6
	`
	result, err := r.db.Exec(
		query,
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
		song.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating song: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no song found with id %s", song.ID)
	}

	return nil
}

func (r *SongRepo) Delete(id string) error {
	query := `DELETE FROM "songs" WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting song: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no song found with id %s", id)
	}

	return nil
}
