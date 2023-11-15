package converter

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
	"github.com/mateusVedoy/go-my-songs.git/src/domain"
)

type Converter struct{}

func (c Converter) ToDomain(dto dto.MusicDTO) (*domain.Music, *domain.BusinessError) {
	music, err := domain.NewMusic(dto.Name, dto.Album, dto.Artist, dto.Duration)

	if err != nil {
		return nil, err
	}

	return music, nil
}

func (c Converter) ToDTO(entity domain.Music) *dto.MusicDTO {
	return &dto.MusicDTO{Name: entity.GetName(), Album: entity.GetAlbum(), Artist: entity.GetArtist(), Duration: entity.GetDuration()}
}

func NewConverterImpl() *Converter {
	return &Converter{}
}
