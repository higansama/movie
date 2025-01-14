package reqres

import (
	"errors"
	"movie-app/internal/utils"
	"movie-app/utils/auth"
)

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type VoteRequest struct {
	Auth    auth.AuthJWT
	Vote    int    `json:"vote"` // [0,1]
	MovieID string `json:"movie_id"`
}

func (v *VoteRequest) Validate() error {
	if v.Vote != 1 && v.Vote != 0 {
		return errors.New("invalid vote value")
	}

	// parse movie id
	_, err := utils.StringToUUID(v.MovieID)
	if err != nil {
		return errors.New("invalid movie id")
	}

	return nil
}
