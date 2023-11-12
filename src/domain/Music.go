package domain

import (
	"errors"
	"regexp"
)

type Music struct {
	name     string
	album    string
	artist   string
	duration string
}

func (music Music) isValid() *BusinessError {
	err := NewBusinessError()

	if isEmpty(music.name) {
		err.AddError(errors.New("Music name cannot be null or empty"))
	}

	if isEmpty(music.album) {
		err.AddError(errors.New("Music album cannot be null or empty"))
	}

	if isEmpty(music.artist) {
		err.AddError(errors.New("Music artist cannot be null or empty"))
	}

	if isEmpty(music.duration) {
		err.AddError(errors.New("Music duration cannot be null or empty"))
	}

	if !isPatternOK(music.duration, "^([0-9]{2}:[0-9]{2})$") {
		err.AddError(errors.New("Music duration is not on pattern desired mm:ss"))
	}

	return err
}

func isPatternOK(value, pattern string) bool {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return regex.MatchString(value)
}

func isEmpty(value string) bool {
	return value == ""
}

func NewMusic(name, album, artist, duration string) (*Music, *BusinessError) {
	music := &Music{name: name, album: album, artist: artist, duration: duration}
	err := music.isValid()

	if err.errors != nil {
		return nil, err
	}

	return music, nil
}

func (music Music) GetName() string {
	return music.name
}

func (music Music) GetAlbum() string {
	return music.album
}

func (music Music) GetArtist() string {
	return music.artist
}

func (music Music) GetDuration() string {
	return music.duration
}
