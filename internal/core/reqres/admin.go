package reqres

import auth "movie-app/utils/auth"

type CreateMovieRequest struct {
	auth.AuthJWT
	Name         string `json:"name" gorm:"<-:create;column:name"`
	Year         int    `json:"year"`
	Description  string `json:"description"`
	Director     string `json:"director" gorm:"<-:create;column:director"`
	CountingView int    `json:"counting_view"`
}
