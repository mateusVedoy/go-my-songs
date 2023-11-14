package usecase

import (
	"fmt"

	"github.com/mateusVedoy/go-my-songs.git/src/application/api"
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
)

func (manager Manager) AddMusic(dto dto.MusicDTO) (*api.ResponseSuccess, *api.ResponseError) {
	music, err := manager.converter.ToDomain(dto)

	if err != nil {
		responseErr := api.SetErrorResponse(400, "There are errors in your request. Check it bellow", err.GetErrors())
		return nil, responseErr
	}

	fmt.Println(music)

	manager.Music = append(manager.Music, *music)

	responseSuccess := api.SetSuccessResponse(201, "Music added successfully!", manager.Music)
	return responseSuccess, nil
}
