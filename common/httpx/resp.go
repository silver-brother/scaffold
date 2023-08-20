package httpx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Resp struct {
	Code int    `json:"code"` // 业务错误码,正常返回为200，错误返回为错误码
	Msg  string `json:"msg"`  // 错误信息，如果状态码为200，msg为OK
	Data any    `json:"data"` // 返回数据
}

func Success(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]string{}
	}
	ctx.Header("trace-id", getTraceIdFromContext(ctx.Request.Context()))
	ctx.Header("span-id", getSpanIdFromContext(ctx.Request.Context()))
	resp := Resp{Code: http.StatusOK, Msg: "OK", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

// Error500 500错误
func Error500(ctx *gin.Context) {
	Error(ctx, http.StatusInternalServerError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]string{})
}

// Error200 200错误
func Error200(ctx *gin.Context, code int, Msg string) {
	Error(ctx, http.StatusOK, code, Msg, map[string]string{})
}

// Error400 40错误
func Error400(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusBadRequest, http.StatusBadRequest, msg, map[string]string{})
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
func Error(ctx *gin.Context, httpCode, code int, Msg string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	ctx.Header("trace-id", getTraceIdFromContext(ctx.Request.Context()))
	ctx.Header("span-id", getSpanIdFromContext(ctx.Request.Context()))
	resp := Resp{Code: code, Msg: Msg, Data: data}
	ctx.JSON(httpCode, resp)
}

func getSpanIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().SpanID().String()
}
func getTraceIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}
