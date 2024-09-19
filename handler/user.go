package handler

import (
	"net/http"

	"github.com/Elys-SaaS/auth/model"
	"github.com/Elys-SaaS/auth/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userService.Create(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (h *Handler) SignIn(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.userService.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *Handler) Refresh(c echo.Context) error {
	req := &refreshTokenRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	claims, err := utils.VerifyJWT(req.RefreshToken, utils.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.AccessForbidden())
	}

	id, ok := (*claims)["id"].(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, utils.AccessForbidden())
	}

	u, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	return c.JSON(http.StatusOK, newUserResponse(u))
}
