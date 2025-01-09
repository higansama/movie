package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID           uuid.UUID  `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name" gorm:"<-:create;column:name"`
	Year         int        `json:"year"`
	Description  string     `json:"description"`
	Slug         string     `json:"slug"`
	Director     string     `json:"director" gorm:"<-:create;column:director"`
	CountingView int        `json:"counting_view"`
	UploadedBy   uuid.UUID  `json:"uploaded_by"`
	UploadedAt   time.Time  `json:"uploaded_at"`
	EditedBy     *uuid.UUID `json:"edited_by"`
	EditedAt     *time.Time `json:"edited_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

// BeforeSave is a GORM hook that is triggered before saving a Movie record.
func (m *Movie) BeforeSave(tx *gorm.DB) (err error) {
	m.Name = strings.ToLower(m.Name)
	m.Director = strings.ToLower(m.Director)
	return
}

// Slug parser
func (m *Movie) SlugMaker() {
	m.Slug = strings.ReplaceAll(strings.ToLower(m.Name), " ", "-")
}

// Add CountingView to the Movie record.
func (m *Movie) AddCountingView() {
	m.CountingView++
}
