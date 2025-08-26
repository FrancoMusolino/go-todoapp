package logger

import (
	"context"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-chi/chi/middleware"
)

type IncomingRequestLog struct {
	Path   string
	Method string
	Query  string
	Body   string
}

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
	color.White("%s\n", msg)
}

func (l *Logger) IncomingRequest(r *http.Request, ctx context.Context) {
	path := r.URL.Path
	method := r.Method

	d := color.New(color.FgCyan, color.Bold)
	d.Printf("[INCOMING REQUEST] %s\n", l.service)

	incomingLog := IncomingRequestLog{
		Path:   path,
		Method: method,
	}

	logReqId(ctx)
	color.White("%+v\n", incomingLog)
}

func logReqId(ctx context.Context) {
	reqID := middleware.GetReqID(ctx)
	color.White("Request ID: %s\n", reqID)
}
