package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"net/http"
)

// @Summary		SignUp
// @Tags			auth
// @Description	create account
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		domain.User	true	"account info"
// @Success		200		{integer}	integer		1
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary		SignIn
// @Tags			auth
// @Description	login
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input	body		signInInput	true	"credentials"
// @Success		200		{string}	string		"token"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.SignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

// @Summary		RefreshToken
// @Tags			auth
// @Description	Refresh token
// @ID				refresh-token
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"token"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/refresh [get]
func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logrus.Infof("%s", cookie)

	accessToken, refreshToken, err := h.services.Authorization.RefreshToken(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}
