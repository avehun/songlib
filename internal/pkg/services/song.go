package services

import (
	"github.com/avehun/songlib/internal/pkg/models"
	"github.com/avehun/songlib/internal/pkg/repo"
	log "github.com/sirupsen/logrus"
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
	songs, err := s.repo.GetAll()
	if err != nil {
		log.Errorf("Unable to get all: %v", err)
	}
	return songs
}
func (s *SongService) RetrieveSong(id string) models.Song {
	song, err := s.repo.GetById(id)
	if err != nil {
		log.Errorf("Unable to get by id: %v", err)
	}
	return song
}
func (s *SongService) DeleteSong(id string) {
	err := s.repo.Delete(id)
	if err != nil {
		log.Errorf("Unable to delete: %v", err)
	}
}
func (s *SongService) ChangeSong(song models.Song) {
	err := s.repo.Update(song)
	if err != nil {
		log.Errorf("Unable to change: %v", err)
	}

}
func (s *SongService) AddSong(song models.Song) {
	err := s.repo.Create(song)
	if err != nil {
		log.Errorf("Unable to add: %v", err)
	}
}
