package handler

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/labstack/echo/v4"
)

type posterCreateRequest struct {
	Poster struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Telephony   string `json:"telephony"`
		Email       string `json:"email"`
		ImageName   string `json:"image_name"`
	} `json:"poster"`
}

type posterUpdateRequest struct {
	Poster struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Telephony   string `json:"telephony"`
		Email       string `json:"email"`
		ImageName   string `json:"image_name"`
	} `json:"poster"`
}

func (request *posterCreateRequest) bind(context echo.Context, poster *model.Poster) error {

	if err := context.Bind(request); err != nil {
		return err
	}

	if err := context.Validate(request); err != nil {
		return err
	}

	poster.Name = request.Poster.Name
	poster.Title = request.Poster.Title
	poster.Description = request.Poster.Description
	poster.Telephony = request.Poster.Telephony
	poster.Email = request.Poster.Email
	poster.ImageName = request.Poster.ImageName

	return nil
}

func (request *posterUpdateRequest) bind(context echo.Context, poster *model.Poster) error {

	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}

	poster.Name = request.Poster.Name
	poster.Title = request.Poster.Title
	poster.Description = request.Poster.Description
	poster.Telephony = request.Poster.Telephony
	poster.Email = request.Poster.Email
	poster.ImageName = request.Poster.ImageName

	return nil
}
