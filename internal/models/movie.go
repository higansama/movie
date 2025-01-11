package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Title       string
	Slug        string
	Director    string
	Description string
	Duration    string
	Casting     []Casting `gorm:"foreignKey:MovieID"`
	Genre       []Genre   `gorm:"many2many:movie_genres"`
	Files       string
	Year        string
	Count       int
	UploadedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
}

// BeforeSave is a GORM hook that is triggered before saving a Movie record.
func (m *Movie) BeforeSave(tx *gorm.DB) (err error) {
	m.Title = strings.ToLower(m.Title)
	m.Director = strings.ToLower(m.Director)
	m.SlugMaker()
	return
}

// Slug parser
func (m *Movie) SlugMaker() {
	m.Slug = strings.ReplaceAll(strings.ToLower(m.Title), " ", "-")
}

// Add CountingView to the Movie record.
func (m *Movie) AddCountingView() {
	m.Count++
}

// get geners ids
func (m *Movie) GetMovieIds() (r []string) {
	for _, v := range m.Genre {
		r = append(r, strconv.Itoa(int(v.ID)))
	}
	return r
}

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name" gorm:"<-:create;column:name"`
	Email    string    `gorm:"unique"`
	Password string
}

type Actor struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`
	Name    string    `json:"name" gorm:"<-:create;column:name"`
	Casting []Casting `gorm:"foreignKey:ActorID"`
}

type Casting struct {
	ID        uuid.UUID `gorm:"primary_key"`
	MovieID   uuid.UUID
	ActorID   uuid.UUID
	Role      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Genre struct {
	ID    uint `gorm:"primaryKey"`
	Title string
	Count int
}

type VotingHistory struct {
	ID          int `gorm:"primaryKey"`
	GenreID     uuid.UUID
	MovieID     uuid.UUID
	UserID      *uuid.UUID
	IpAddress   string
	IsLike      bool
	DateCreated time.Time `gorm:"autoCreateTime"`
}

type WathcingHistory struct {
	ID      int `gorm:"primaryKey"`
	UserID  *uuid.UUID
	MovieID uuid.UUID
	WatchAt time.Time `gorm:"autoCreateTime"`
}
