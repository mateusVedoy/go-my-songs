package api

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
)

type Error struct {
	Errors []string `json:"errors"`
}

type ResponseSuccess struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Content []dto.MusicDTO `json:"result"`
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Content Error  `json:"result"`
}

func buildError(err []error) *Error {
	Error := &Error{}
	for _, e := range err {
		Error.Errors = append(Error.Errors, e.Error())
	}
	return Error
}

func SetErrorResponse(status int, message string, content []error) *ResponseError {
	err := buildError(content)
	responseError := ResponseError{Status: status, Message: message, Content: *err}
	return &responseError
}

func SetSuccessResponse(status int, message string, content []dto.MusicDTO) *ResponseSuccess {
	success := ResponseSuccess{Status: status, Message: message, Content: content}
	return &success
}
