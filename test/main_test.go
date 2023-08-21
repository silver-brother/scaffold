package test

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"scaffold/internal/config"
	"scaffold/internal/router"
	"scaffold/internal/service"
	"scaffold/pkg/conf"
	"scaffold/pkg/validator"
	"testing"
)

var configFile = flag.String("f", "../config/config.yaml", "the config file")

var e *gin.Engine

func TestMain(m *testing.M) {
	flag.Parse()

	// 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	c.Log.Dir = "../logs"
	c.Mode = gin.ReleaseMode

	svcCtx := service.NewServiceContext(c)

	// 自定义验证器
	validator.CustomValidator()

	gin.SetMode(c.Mode)
	e = gin.Default()
	router.Router(e, svcCtx)

	code := m.Run()
	// 退出
	os.Exit(code)
}
