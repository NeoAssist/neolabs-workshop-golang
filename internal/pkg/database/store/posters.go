package store

import (
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/jinzhu/gorm"
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

// GetByID .
func (as *PosterStore) GetByID(id string) (*model.Poster, error) {
	var m model.Poster
	if err := as.db.Where(&model.Poster{ID: id}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetByEmail .
func (as *PosterStore) GetByEmail(e string) (*model.Poster, error) {
	var m model.Poster
	if err := as.db.Where(&model.Poster{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// GetByName .
func (as *PosterStore) GetByName(name string) (*model.Poster, error) {
	var m model.Poster
	if err := as.db.Where(&model.Poster{Name: name}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// GetByDescription .
func (as *PosterStore) GetByDescription(description string) (*model.Poster, error) {
	var m model.Poster
	if err := as.db.Where(&model.Poster{Description: description}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// GetByTitle .
func (as *PosterStore) GetByTitle(title string) (*model.Poster, error) {
	var m model.Poster
	if err := as.db.Where(&model.Poster{Title: title}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// Create .
func (as *PosterStore) Create(u *model.Poster) (err error) {
	transaction := as.db.Begin()

	if err := transaction.Model(u).Create(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}

// Update .
func (as *PosterStore) Update(u *model.Poster) error {
	transaction := as.db.Begin()

	if err := transaction.Model(u).Update(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}
