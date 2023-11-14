package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
	usecase "github.com/mateusVedoy/go-my-songs.git/src/application/useCase"
)

type Controller struct {
	manager usecase.Manager
}

func (c Controller) Add(writer http.ResponseWriter, request *http.Request) {
	dto := &dto.MusicDTO{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		http.Error(writer, "request failed to parse JSON body. Check your request!", http.StatusBadRequest)
	}
	successResp, errorResp := c.manager.AddMusic(*dto)

	if errorResp != nil {
		writer.WriteHeader(errorResp.Status)
		writer.Write([]byte(errorResp.Message))
	}

	writer.WriteHeader(successResp.Status)
	writer.Write([]byte(successResp.Message))
}

func NewController(manager *usecase.Manager) *Controller {
	return &Controller{manager: *manager}
}
