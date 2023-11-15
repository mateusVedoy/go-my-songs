package api

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
)

type ResponseSuccess struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Content []dto.MusicDTO `json:"content"`
}

type ResponseError struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Content []error `json:"content"`
}

func SetErrorResponse(status int, message string, content []error) *ResponseError {
	err := ResponseError{Status: status, Message: message, Content: content}
	return &err
}

func SetSuccessResponse(status int, message string, content []dto.MusicDTO) *ResponseSuccess {
	success := ResponseSuccess{Status: status, Message: message, Content: content}
	return &success
}
