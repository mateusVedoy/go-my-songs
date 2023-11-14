package dto

type MusicDTO struct {
	Name     string `json:name`
	Album    string `json:album`
	Artist   string `json:artist`
	Duration string `json:duration`
}
