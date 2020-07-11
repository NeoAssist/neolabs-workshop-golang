package store

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// PosterStore .
type PosterStore struct {
	db *gorm.DB
}

// NewPosterStore .
func NewPosterStore(db *gorm.DB) *PosterStore {
	return &PosterStore{
		db: db,
	}
}

func (ps *PosterStore) GetAll(offset int, limit int) ([]model.Poster, int, error) {
	var (
		posters []model.Poster
		count   int
	)

	ps.db.Model(&posters).Count(&count)
	ps.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&posters)

	return posters, count, nil
}

// GetByID .
func (ps *PosterStore) GetByID(id uuid.UUID) (*model.Poster, error) {
	var m model.Poster

	if err := ps.db.Where(&model.Poster{ID: id}).Find(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// GetByTitle .
func (ps *PosterStore) GetByTitle(title string) ([]model.Poster, int, error) {
	var (
		posters []model.Poster
		count   int
	)

	ps.db.Model(&posters)
	ps.db.Where("title LIKE ?", "%"+title+"%").Find(&posters).Count(&count)

	return posters, count, nil
}

// Create .
func (ps *PosterStore) Create(u *model.Poster) (err error) {
	transaction := ps.db.Begin()

	if err := transaction.Model(u).Create(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}

// Update .
func (ps *PosterStore) Update(u *model.Poster) error {
	transaction := ps.db.Begin()

	if err := transaction.Model(u).Update(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}
