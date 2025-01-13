package reqres

import (
	"movie-app/utils/auth"
	"strings"
)

type CreateMovieRequest struct {
	auth.AuthJWT
	Movie MovieCreate `json:"movie"`
	Actor []Actor     `json:"actor"`
}

func (c CreateMovieRequest) JoinTheGenre() string {
	return strings.Join(c.Movie.Genres, ",")
}

type MovieCreate struct {
	ID          string   `json:"string"`
	Title       string   `json:"title"`
	Director    string   `json:"director"`
	Description string   `json:"description"`
	Duration    string   `json:"duration"`
	Genres      []string `json:"genre"`
	Files       string   `json:"files"`
	Year        string   `json:"year"`
}

type Actor struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}
type CreateMovieForm struct {
	Title       string      `form:"title" binding:"omitempty"`
	Director    string      `form:"director" binding:"omitempty"`
	Description string      `form:"description" binding:"omitempty"`
	Duration    string      `form:"duration" binding:"omitempty"`
	Genres      string      `form:"genres" binding:"omitempty"`
	Files       string      `form:"files" binding:"omitempty"`
	Year        string      `form:"year" binding:"omitempty"`
	Actors      []ActorForm `form:"actors" binding:"omitempty,dive"`
}

type ActorForm struct {
	ID   string `form:"id" binding:"required"`
	Role string `form:"role" binding:"required"`
}

type EditMovieRequest struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Genres      string `json:"genres"`
	Year        string `json:"year"`
}
