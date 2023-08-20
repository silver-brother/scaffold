package httpx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]string{}
	}
	ctx.Header("trace-id", getTraceIdFromContext(ctx.Request.Context()))
	ctx.Header("span-id", getSpanIdFromContext(ctx.Request.Context()))
	resp := Resp{Code: http.StatusOK, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

// Error500 500错误
func Error500(ctx *gin.Context) {
	Error(ctx, http.StatusInternalServerError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]string{})
}

// Error200 200错误
func Error200(ctx *gin.Context, code int, message string) {
	Error(ctx, http.StatusOK, code, message, map[string]string{})
}

// Error400 40错误
func Error400(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusBadRequest, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), msg)
}

// Error401 401错误
func Error401(ctx *gin.Context) {
	Error(ctx, http.StatusUnauthorized, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]string{})
}

// Error404 404错误
func Error404(ctx *gin.Context) {
	Error(ctx, http.StatusNotFound, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]string{})
}

// Error 通用错误
func Error(ctx *gin.Context, httpCode, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	ctx.Header("trace-id", getTraceIdFromContext(ctx.Request.Context()))
	ctx.Header("span-id", getSpanIdFromContext(ctx.Request.Context()))
	resp := Resp{Code: code, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}

func getSpanIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().SpanID().String()
}
func getTraceIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}
