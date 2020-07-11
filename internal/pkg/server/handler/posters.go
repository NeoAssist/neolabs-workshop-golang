package handler

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/errors"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (h *Handler) UploadPosterImage(context echo.Context) error {
	file, err := context.FormFile("file")

	if err != nil {
		return err
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	// Destination
	staticFolder := filepath.Join(utils.GetEnv("STATIC_FOLDER"), file.Filename)
	dst, err := os.Create(staticFolder)

	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, newPosterCreateResponse("Image saved successfully."))
}

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

func (h *Handler) GetAllPosters(context echo.Context) error {
	var (
		result []model.Poster
		count  int
	)

	offset, err := strconv.Atoi(context.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(context.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	result, count, err = h.posterStore.GetAll(offset, limit)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	if count == 0 {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterListResponse(result, count))
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

func (h *Handler) GetPosterByTitle(context echo.Context) error {
	posterTitle := context.QueryParam("query")

	result, count, err := h.posterStore.GetByTitle(posterTitle)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if result == nil {
		return context.JSON(http.StatusNotFound, errors.NotFound())
	}

	return context.JSON(http.StatusOK, newPosterListResponse(result, count))
}
