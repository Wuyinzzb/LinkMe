package ioc

import (
	"LinkMe/internal/api"
	"github.com/gin-gonic/gin"
)

// InitWebServer 初始化web服务
func InitWebServer(userHdl *api.UserHandler, postHdl *api.PostHandler, historyHdl *api.HistoryHandler, checkHdl *api.CheckHandler, m []gin.HandlerFunc, permHdl *api.PermissionHandler) *gin.Engine {
	server := gin.Default()
	server.Use(m...)
	userHdl.RegisterRoutes(server)
	postHdl.RegisterRoutes(server)
	historyHdl.RegisterRoutes(server)
	checkHdl.RegisterRoutes(server)
	permHdl.RegisterRoutes(server)
	return server
}
