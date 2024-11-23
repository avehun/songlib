package services

import "github.com/avehun/songlib/internal/pkg/models"

func ListSongs() []models.Song {
	println("list songs service")
	return []models.Song{}
}
func RetrieveSong(id string) models.Song {
	println("retrieve songs service")
	return models.Song{}
}
func DeleteSong(id string) {
	println("delete song service")
}
func ChangeSong(id string) {
	println("change song service")
}
func AddSong(models.Song) {
	println("add song service")
}
