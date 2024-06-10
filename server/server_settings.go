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
