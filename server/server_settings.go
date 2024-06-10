package server

import (
	"all_inclusive_app/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine  *gin.Engine
	Service *service.Service
}

func NewHandler(service *service.Service, engine *gin.Engine) *Handler {
	return &Handler{Service: service, Engine: engine}
}

func (h Handler) AllRoutes() {
	h.Engine.POST("/login", h.Login)
	h.Engine.POST("/register", h.Register)
	h.Engine.GET("/get_all_users", h.GetAllUsers)
	h.Engine.GET("/get_user/:id", h.GetUser)
	h.Engine.PUT("/update_user/:id", h.UpdateUser)
	h.Engine.DELETE("/delete_user/:id", h.DeleteUser)

}
