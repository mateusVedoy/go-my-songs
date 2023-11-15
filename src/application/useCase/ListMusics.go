package usecase

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/api"
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
	"github.com/mateusVedoy/go-my-songs.git/src/infrastructure/database"
)

func (manager Manager) ListAll() (*api.ResponseSuccess, *api.ResponseError) {
	var dtoList []dto.MusicDTO
	for _, value := range database.DB {
		dtoList = append(dtoList, *manager.converter.ToDTO(value))
	}

	success := api.SetSuccessResponse(200, "Music fetched bellow", dtoList)

	return success, nil
}
