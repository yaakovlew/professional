package handler

import (
	"github.com/gin-gonic/gin"
	"proffesional/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/group", h.AddGroup)
	router.GET("groups", h.GetAllGroups)
	router.GET("/group/:id", h.GetGroupForId)
	router.PUT("/group/:id", h.ChangeSomeGroup)
	router.DELETE("/group/:id", h.DeleteGroupForId)
	return router
}
