package server

import (
	"all_inclusive_app/models"
	"github.com/sirupsen/logrus"
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
	var loginData models.LoginStruct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, please, try again"})
		logrus.Println("Cannot bind JSON in func Login, Error:", err.Error())
		return
	}
	user, err := h.Service.Login(loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, cannot find user with this parameters"})
		logrus.Println(err.Error())
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, cannot generate token for this user"})
		logrus.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{Token: token})
}
