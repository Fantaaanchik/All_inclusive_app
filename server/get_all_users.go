package server

import (
	"all_inclusive_app/models"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce  json
// @Success 200 {array} models.User
// @Failure 401 {object} models.ErrorResponse
// @Router /get_all_users [get]

// GetAllUsers - This function retrieves all users from the database.
func (h Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{ErrorDescription: "cannot get all users"})
		return
	}

	c.JSON(http.StatusOK, models.GetAllUsersResponse{Users: users})
}
