package handler

import (
	"github.com/labstack/echo/v4"
)

// Register .
func (h *Handler) Register(v1 *echo.Group) {

	posters := v1.Group("/poster")

	posters.POST("", h.CreatePoster)
	posters.GET("/:id", h.GetPosterById)
	posters.GET("/:name", h.GetPosterByName)
	posters.GET("/:description", h.GetPosterByDescription)
	posters.GET("/:email", h.GetPosterByEmail)
	posters.PUT("/:id", h.UpdatePoster)
}
