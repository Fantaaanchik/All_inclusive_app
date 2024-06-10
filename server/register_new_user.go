package server

import (
	"all_inclusive_app/models"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.RegisterStruct true "Add new user"
// @Success 200 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /register [post]

// Register - Registering a new user in the DB
func (h Handler) Register(c *gin.Context) {
	var user models.RegisterStruct

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "something went`s wrong"})
		logrus.Println("Cannot bind JSON in func Register, Error:", err.Error())
		return
	}

	err := h.Service.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, try again"})
		logrus.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, models.CreateUserResponse{Message: "user successfully registered"})
}
