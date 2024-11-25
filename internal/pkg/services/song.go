package services

import (
	"fmt"

	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/avehun/songlib/internal/pkg/repo"
)

type SongService struct {
	repo *repo.SongRepo
}

func NewSongService(repo *repo.SongRepo) *SongService {
	return &SongService{
		repo: repo,
	}
}
func (s *SongService) ListSongs() []models.Song {
	println("list songs service")
	return []models.Song{}
}
func (s *SongService) RetrieveSong(id string) models.Song {
	println("retrieve songs service")
	song, err := s.repo.GetById(id)
	if err != nil {
		fmt.Errorf("Unable to get by id: %v", err)
	}
	return *song
}
func (s *SongService) DeleteSong(id string) {
	println("delete song service")
}
func (s *SongService) ChangeSong(id string) {
	println("change song service")
}
func (s *SongService) AddSong(models.Song) {
	println("add song service")
}
