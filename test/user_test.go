package test

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"scaffold/internal/config"
	"scaffold/internal/router"
	"scaffold/internal/service"
	"scaffold/pkg/conf"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var configFile = flag.String("f", "../config/config.yaml", "the config file")

func TestUserListRouter(t *testing.T) {
	flag.Parse()

	// 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := service.NewServiceContext(c)

	gin.SetMode(c.Mode)
	e := gin.Default()

	router.Router(e, svcCtx)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/list", nil)
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
