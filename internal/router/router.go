package router

import (
	"net/http"
	"scaffold/internal/handler"
	"scaffold/internal/middleware"
	"scaffold/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine, svcCtx *service.ServiceContext) {
	// middlewares
	r.Use(middleware.TraceMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	// swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// user handler
	userHandle := handler.NewUserHandle(svcCtx)
	userGroup := r.Group("user")
	{
		userGroup.GET("/", userHandle.List)
		userGroup.POST("/", userHandle.Create)
	}
}
