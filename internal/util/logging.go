package util

import (
	"context"
	"log"
	"log/slog"
	"os"
)

type levelHandler struct {
	handler slog.Handler
}

func NewLevelHandler(h slog.Handler) *levelHandler {
	return &levelHandler{h}
}

func (h *levelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= slog.LevelDebug
}

func (h *levelHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

func (h *levelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.handler.WithAttrs(attrs)
}

func (h *levelHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}

func ConfigureLogging() {

	handler := &levelHandler{
		handler: slog.Default().Handler(),
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmsgprefix | log.LstdFlags)
}
