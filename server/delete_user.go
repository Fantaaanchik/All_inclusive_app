package server

import (
	"all_inclusive_app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Description Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} models.DeleteUserResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /delete_user/{id} [delete]

// DeleteUser - Delete user in DB with ID from context parameter
func (h Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteUser(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorDescription: "bad request"})
		return
	}
	c.JSON(http.StatusOK, models.DeleteUserResponse{Message: "user successfully deleted!"})
}
