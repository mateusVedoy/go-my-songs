package dto

type MusicDTO struct {
	name     string `json:name`
	album    string `json:album`
	artist   string `json:artist`
	duration string `json:duration`
}

func (dto MusicDTO) GetName() string {
	return dto.name
}

func (dto MusicDTO) GetAlbum() string {
	return dto.album
}

func (dto MusicDTO) GetArtist() string {
	return dto.artist
}

func (dto MusicDTO) GetDuration() string {
	return dto.duration
}

func NewMusicDTO(name, album, artist, duration string) *MusicDTO {
	return &MusicDTO{name: name, album: album, artist: artist, duration: duration}
}
