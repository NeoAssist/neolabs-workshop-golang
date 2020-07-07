package handler

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// CreatePoster .
func (h *Handler) CreatePoster(context echo.Context) error {
	var account model.Poster

	req := &posterCreateRequest{}

	if err := req.bind(context, &account); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	err := h.posterStore.Create(&account)

	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	return context.JSON(http.StatusCreated, newPosterCreateResponse("Registry created successfully."))
}

// UpdatePoster .
func (h *Handler) UpdatePoster(context echo.Context) error {
	id, err := uuid.FromString(context.Param("id"))
	poster, err := h.posterStore.GetByID(id)

	req := &posterUpdateRequest{}

	if err := req.bind(context, poster); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	h.posterStore.Update(poster)

	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	return context.JSON(http.StatusNoContent, newPosterUpdateResponse("Registry updated successfully."))
}

func (h *Handler) GetPosterById(context echo.Context) error {
	posterId, err := uuid.FromString(context.Param("id"))

	result, err := h.posterStore.GetByID(posterId)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterResponse(result))
}

func (h *Handler) GetPosterByName(context echo.Context) error {
	posterName := context.Param("name")

	result, err := h.posterStore.GetByName(posterName)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterResponse(result))
}

func (h *Handler) GetPosterByDescription(context echo.Context) error {
	posterDescription := context.Param("description")

	result, err := h.posterStore.GetByDescription(posterDescription)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterResponse(result))
}

func (h *Handler) GetPosterByEmail(context echo.Context) error {
	posterDescription := context.Param("email")

	result, err := h.posterStore.GetByEmail(posterDescription)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterResponse(result))
}
