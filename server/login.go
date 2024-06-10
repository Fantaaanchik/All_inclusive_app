package server

import (
	"all_inclusive_app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description Login with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.LoginStruct true "Login User"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /login [post]

// Login - Sign in with login and password
func (h Handler) Login(c *gin.Context) {
	var user models.LoginStruct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, please, try again"})
		return
	}
	token, err := h.Service.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, cannot find user with this parameters"})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{Token: token})
}
