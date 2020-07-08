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
	posters.GET("/search/:id", h.GetPosterById)
	posters.GET("/search/name", h.GetPosterByName)
	posters.GET("/search/title", h.GetPosterByTitle)
	posters.GET("/search/description", h.GetPosterByDescription)
	posters.GET("/search/email", h.GetPosterByEmail)
}
