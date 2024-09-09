package service

import "github.com/gin-gonic/gin"

type RestHandler interface {
	RegisterHandlerRoutes(r *gin.RouterGroup)
}
