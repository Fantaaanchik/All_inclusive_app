package server

import (
	"all_inclusive_app/models"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Update user
// @Description Update user with id and some parameters
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User Parameters"
// @Success 200 {object} models.UpdateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /update_user/{id} [put]

// UpdateUser - Updating the user parameters with the ID and new parameters
func (h Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUserStruct

	id := c.Query("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, wrong parameters"})
		logrus.Println("Find Error in func UpdateUser on path server, Error description:", err.Error())
		return
	}

	err := h.Service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{ErrorDescription: "cannot update user with this parameters"})
		logrus.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, models.UpdateUserResponse{Message: "user successfully updated"})
}
