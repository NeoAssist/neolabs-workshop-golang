package handler

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"time"
)

type posterResponse struct {
	Poster struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"poster"`
}

type singlePosterResponse struct {
	Poster struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Email       string    `json:"description"`
		Telephony   string    `json:"telephony"`
		Title       string    `json:"title"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
}

type allPostersResponse struct {
	Posters []*posterResponse `json:"articles"`
}

func newPosterCreateResponse(message string) *posterResponse {
	response := new(posterResponse)

	response.Poster.Status = "created"
	response.Poster.Message = message

	return response
}

func newPosterUpdateResponse(message string) *posterResponse {
	response := new(posterResponse)

	response.Poster.Status = "updated"
	response.Poster.Message = message

	return response
}

func newPosterResponse(poster *model.Poster) *singlePosterResponse {
	response := new(singlePosterResponse)

	response.Poster.Name = poster.Name
	response.Poster.Description = poster.Description
	response.Poster.Email = poster.Email
	response.Poster.Telephony = poster.Telephony
	response.Poster.Title = poster.Title
	response.Poster.CreatedAt = poster.CreatedAt
	response.Poster.UpdatedAt = poster.UpdatedAt

	return response
}
