package api

import "github.com/mateusVedoy/go-my-songs.git/src/domain"

type ResponseSuccess struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Content []domain.Music `json:"content"`
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

func SetSuccessResponse(status int, message string, content []domain.Music) *ResponseSuccess {
	success := ResponseSuccess{Status: status, Message: message, Content: content}
	return &success
}
