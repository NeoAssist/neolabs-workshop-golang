package handler

import (
	interfaces "github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/interface"
)

// Handler .
type Handler struct {
	posterStore interfaces.Store
}

// NewHandler Handle the requests
func NewHandler(as interfaces.Store) *Handler {
	return &Handler{
		posterStore: as,
	}
}
