package usecase

import "github.com/mateusVedoy/go-my-songs.git/src/domain"

type Manager struct {
	converter domain.IConvert
}

func NewManager(converter domain.IConvert) *Manager {
	return &Manager{converter: converter}
}
