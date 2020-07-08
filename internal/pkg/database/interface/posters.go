package interfaces

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	uuid "github.com/satori/go.uuid"
)

// Store .
type PosterStore interface {
	GetAll(offset int, limit int) ([]model.Poster, int, error)
	GetByID(uuid.UUID) (*model.Poster, error)
	GetByEmail(string) (*model.Poster, error)
	GetByName(string) (*model.Poster, error)
	GetByTitle(string) (*model.Poster, error)
	GetByDescription(string) (*model.Poster, error)
	Create(*model.Poster) error
	Update(*model.Poster) error
}
