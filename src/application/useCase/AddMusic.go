package usecase

import (
	"fmt"

	"github.com/mateusVedoy/go-my-songs.git/src/application/api"
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
	"github.com/mateusVedoy/go-my-songs.git/src/infrastructure/database"
)

func (manager Manager) AddMusic(musicDTO dto.MusicDTO) (*api.ResponseSuccess, *api.ResponseError) {
	music, err := manager.converter.ToDomain(musicDTO)

	if err != nil {
		responseErr := api.SetErrorResponse(400, "There are errors in your request. Check it bellow", err.GetErrors())
		return nil, responseErr
	}

	database.DB = append(database.DB, *music)
	fmt.Println("Music saved!")

	var dtoArr []dto.MusicDTO
	dtoArr = append(dtoArr, musicDTO)

	responseSuccess := api.SetSuccessResponse(201, "Music added successfully!", dtoArr)
	return responseSuccess, nil
}
