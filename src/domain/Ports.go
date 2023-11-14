package domain

import "github.com/mateusVedoy/go-my-songs.git/src/application/dto"

type IConvert interface {
	ToDomain(dto dto.MusicDTO) (*Music, *BusinessError)
	ToDTO(entity Music) *dto.MusicDTO
}
