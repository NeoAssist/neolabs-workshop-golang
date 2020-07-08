package handler

import (
	interfaces "github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/interface"
)

// Handler .
type Handler struct {
	posterStore interfaces.PosterStore
}

// NewHandler Handle the requests
func NewHandler(ps interfaces.PosterStore) *Handler {
	return &Handler{
		posterStore: ps,
	}
}
