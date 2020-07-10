package handler

import (
	"github.com/labstack/echo/v4"
)

// Register .
func (h *Handler) Register(v1 *echo.Group) {

	posters := v1.Group("/poster")

	posters.POST("", h.CreatePoster)
	posters.POST("/upload", h.UploadPosterImage)

	posters.PUT("/:id", h.UpdatePoster)

	posters.GET("", h.GetAllPosters)
	posters.GET("/:id", h.GetPosterById)
	posters.GET("/title", h.GetPosterByTitle)
}
