package reqres

import (
	auth "movie-app/utils/auth"
)

type CreateMovieRequest struct {
	auth.AuthJWT
	Movie MovieCreate `json:"movie"`
	Actor []Actor     `json:"actor"`
}

type MovieCreate struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Genres      string `json:"genre"`
	Files       string `json:"files"`
	Year        string `json:"year"`
}

type Actor struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type EditMovieRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Director    string `json:"director"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Genres      string `json:"genres"`
	Year        string `json:"year"`
}
