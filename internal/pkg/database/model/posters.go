package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm" // not blank
	uuid "github.com/satori/go.uuid"
)

// Account model
type Poster struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `gorm:"size:100"`
	Title       string    `gorm:"size:30"`
	Description string    `gorm:"size:255"`
	Telephony   string    `gorm:"size:20"`
	Email       string    `gorm:"size:30"`
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
