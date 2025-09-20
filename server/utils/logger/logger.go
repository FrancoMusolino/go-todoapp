package logger

import (
	"context"

	"github.com/fatih/color"
	"github.com/go-chi/chi/middleware"
)

type Logger struct {
	service string
}

func NewLogger(service string) *Logger {
	return &Logger{
		service: service,
	}
}

func (l *Logger) Info(ctx context.Context, stack, msg string) {
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("[INFO] %s - %s\n", l.service, stack)

	logReqId(ctx)
	color.White("%s\n", msg)
}

func (l *Logger) Error(ctx context.Context, stack, msg string) {
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("[INFO] %s - %s\n", l.service, stack)

	logReqId(ctx)
	color.Red("%s\n", msg)
}

func logReqId(ctx context.Context) {
	reqID := middleware.GetReqID(ctx)
	color.White("Request ID: %s\n", reqID)
}
