package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserListRouter(t *testing.T) {
	params := map[string]interface{}{
		// "age":      "10",
		// "birthday": "2020-0101",
		// "email":    "zhangsanfeng@163.com",
		// "mobile":   "138888",
		// "nickname": "张三疯",
		"username": "张三",
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"code":404,"msg":"Not Found","data":{}}`, w.Body.String())
}
