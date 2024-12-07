package main

import (
	"context"
	"errors"
	"example/internal/library"
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	slog.Debug("Debug Level", "MyKey", "Detailed information for debugging")
	slog.Info("Info Level", "AnotherKey", "General operational information")
	slog.Warn("Warn Level", "ImportantKey", "Potential issue detected")
	slog.Error("Error Level", "ErrorKey", "A serious issue occurred")
	slog.Error("Bad string", "No key provided") // go vet
	slog.Error("Nice", slog.Int("user_id", 123))
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"My request",
		slog.String("method", "GET"),
		slog.Int("request_processing_ms", 132),
		slog.String("path", "/hello/world"),
		slog.Int("status", 200),
	)

	logger2 := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger2.Info("TextHandler Example", "Content", "Logging in text format")

	loggerInfoOnly := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	loggerInfoOnly.Info("Will be printed")
	loggerInfoOnly.Debug("Will be skipped")

	buildInfo, _ := debug.ReadBuildInfo()
	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
	child.Info("Hello", "Content", "Golang slog is best for structured logging")

	group := slog.Group("Usage",
		slog.String("Item1", "random string"),
		slog.Int("Item2", 123),
		slog.Int("Item3", 100500),
	)
	child.Info("slog Groups", group)

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug, // or get level from env var
	}
	loggerWithSource := slog.New(slog.NewJSONHandler(os.Stderr, opts))
	library.Call(loggerWithSource)

	h := &ContextHandler{slog.NewJSONHandler(os.Stdout, opts)}
	logger4 := slog.New(h)
	ctx := AppendCtx(context.Background(), slog.String("request_id", "req-123"))
	ctx2 := AppendCtx(ctx, slog.String("user_id", "user-123"))
	logger4.InfoContext(ctx2, "Request processed", slog.String("param", "value"))

	logger4.ErrorContext(ctx, "Request failed", slog.Any("error", errors.New("something bad")))

}

type ctxKey string

const (
	slogAttrs ctxKey = "slog_attrs"
)

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogAttrs).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds an slog attribute to the provided context so that it will be
// included in any Record created with such context
func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogAttrs).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogAttrs, v)
	}

	var v []slog.Attr
	v = append(v, attr)
	return context.WithValue(parent, slogAttrs, v)
}
