package reqres

import "time"

type MovieResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Director    string    `json:"director"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"`
	Artist      []string  `json:"artist"`
	Genres      string    `json:"genres"`
	Files       string    `json:"files"`
	Year        string    `json:"year"`
	Count       int       `json:"count"`
	UploadedAt  time.Time `json:"uploaded_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
