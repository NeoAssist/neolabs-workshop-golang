package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm" // not blank
	uuid "github.com/satori/go.uuid"
)

// Account model
type Poster struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Name        string    `gorm:"size:100";not null`
	Title       string    `gorm:"size:30";not null;index:title_idx`
	Description string    `gorm:"size:255";not null`
	Telephony   string    `gorm:"size:20";not null`
	Email       string    `gorm:"size:30";not null`
	ImageName   string    `gorm:"size:100";not null`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (account *Poster) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)

	return nil
}
