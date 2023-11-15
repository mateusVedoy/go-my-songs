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

	writer.Header().Set("Content-Type", "application/json")

	if errorResp != nil {
		resp, _ := json.Marshal(errorResp)
		writer.Write([]byte(resp))
	} else {
		resp, _ := json.Marshal(successResp)
		writer.Write([]byte(resp))
	}
}

func (c Controller) ListAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	successResp, errorResp := c.manager.ListAll()
	if errorResp != nil {
		resp, _ := json.Marshal(errorResp)
		writer.Write([]byte(resp))
	} else {
		resp, _ := json.Marshal(successResp)
		writer.Write([]byte(resp))
	}

}

func NewController(manager *usecase.Manager) *Controller {
	return &Controller{manager: *manager}
}
