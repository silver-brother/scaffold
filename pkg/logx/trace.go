package logx

import (
	"context"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId := getSpanIdFromContext(ctx) // as per your tracing framework
	traceId := getTraceIdFromContext(ctx)

	e.Str("span-id", spanId)
	e.Str("trace-id", traceId)
}

func getSpanIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().SpanID().String()
}
func getTraceIdFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}
