package server

import (
	"all_inclusive_app/models"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /get_user/{id} [get]

// GetUser - This function retrieves a user from the database using the identifier obtained from the context parameter
func (h Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, cannot find user with this ID"})
		logrus.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, models.GetUserResponse{User: user})
}
