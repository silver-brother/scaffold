package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func PostJsonQuest(e *gin.Engine, uri string, params map[string]any) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(params)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonByte))
	e.ServeHTTP(w, req)

	return w
}

func GetFormQuery(e *gin.Engine, uri string, params map[string]any) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()
	e.ServeHTTP(w, req)

	return w
}
