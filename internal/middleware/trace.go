package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		otp := otel.GetTracerProvider()
		tp := trace.NewTracerProvider(trace.WithSampler(trace.AlwaysSample()))
		otel.SetTracerProvider(tp)
		defer otel.SetTracerProvider(otp)

		var opts []oteltrace.SpanStartOption
		ctx1, span := tp.Tracer("trace-id").Start(c.Request.Context(), "span-id", opts...)
		c.Request = c.Request.WithContext(ctx1)
		defer span.End()
	}
}
