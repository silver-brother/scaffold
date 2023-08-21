package router

import (
	"github.com/gin-gonic/gin"
	"scaffold/internal/handler"
	"scaffold/internal/service"
)

func V1Router(v1 *gin.RouterGroup, svcCtx *service.ServiceContext) {

	// user handler
	userHandle := handler.NewUserHandle(svcCtx)
	userGroup := v1.Group("user")
	{
		userGroup.GET("", userHandle.List)
		userGroup.POST("", userHandle.Create)
	}
}
