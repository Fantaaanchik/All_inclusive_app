package server

import (
	"all_inclusive_app/models"
	"log"
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
func (h Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUserStruct

	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request, wrong parameters"})
		log.Println("Find Error in func UpdateUser on path server, Error description:", err.Error())
		return
	}

	err := h.Service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{ErrorDescription: "cannot update user with this parameters"})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, models.UpdateUserResponse{Message: "user successfully updated"})
}
