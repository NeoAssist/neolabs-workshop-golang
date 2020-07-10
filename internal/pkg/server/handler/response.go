package handler

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	uuid "github.com/satori/go.uuid"
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
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Email       string    `json:"email"`
		Telephony   string    `json:"telephony"`
		Title       string    `json:"title"`
		ImageName   string    `json:"image_name"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	} `json:"poster"`
}

type allPostersResponse struct {
	Posters      []*singlePosterResponse `json:"posters"`
	PostersCount int                     `json:"postersCount"`
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

	response.Poster.ID = poster.ID
	response.Poster.Name = poster.Name
	response.Poster.Description = poster.Description
	response.Poster.Email = poster.Email
	response.Poster.Telephony = poster.Telephony
	response.Poster.Title = poster.Title
	response.Poster.ImageName = poster.ImageName
	response.Poster.CreatedAt = poster.CreatedAt
	response.Poster.UpdatedAt = poster.UpdatedAt

	return response
}

func newPosterListResponse(posters []model.Poster, count int) *allPostersResponse {
	response := new(allPostersResponse)
	response.Posters = make([]*singlePosterResponse, 0)

	for _, value := range posters {
		pr := new(singlePosterResponse)

		pr.Poster.ID = value.ID
		pr.Poster.Name = value.Name
		pr.Poster.Description = value.Description
		pr.Poster.Email = value.Email
		pr.Poster.Telephony = value.Telephony
		pr.Poster.Title = value.Title
		pr.Poster.CreatedAt = value.CreatedAt
		pr.Poster.UpdatedAt = value.UpdatedAt
		pr.Poster.ImageName = value.ImageName

		response.Posters = append(response.Posters, pr)
	}

	response.PostersCount = count

	return response
}
