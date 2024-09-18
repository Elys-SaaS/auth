package handler

import (
	"github.com/Elys-SaaS/auth/model"
	"github.com/Elys-SaaS/auth/utils"
)

type userResponse struct {
	User struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.AccessToken = utils.GenerateJWT(u.ID, utils.AccessToken)
	r.User.RefreshToken = utils.GenerateJWT(u.ID, utils.RefreshToken)
	return r
}
