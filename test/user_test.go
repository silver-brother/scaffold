package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserList(t *testing.T) {
	w := userListQuest()

	assert.Equal(t, 200, w.Code)
	code := gjson.Get(w.Body.String(), "code").Int()
	assert.Equal(t, 200, int(code))
}

func BenchmarkUserList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := userListQuest()

		assert.Equal(b, 200, w.Code)
		code := gjson.Get(w.Body.String(), "code").Int()
		assert.Equal(b, 200, int(code))
	}
}

func userListQuest() *httptest.ResponseRecorder {
	params := map[string]any{
		"page":      "1",
		"page_size": "10",
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user", nil)

	var query strings.Builder
	for k, v := range params {
		query.WriteString(fmt.Sprintf("%s=%s&", k, v))
	}
	req.URL.RawQuery = strings.TrimRight(query.String(), "&")
	e.ServeHTTP(w, req)
	return w
}
