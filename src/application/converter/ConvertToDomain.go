package converter

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
	"github.com/mateusVedoy/go-my-songs.git/src/domain"
)

func Convert(dto dto.MusicDTO) (*domain.Music, *domain.BusinessError) {
	music, err := domain.NewMusic(dto.GetName(), dto.GetAlbum(), dto.GetArtist(), dto.GetDuration())

	if err != nil {
		return nil, err
	}

	return music, nil
}
