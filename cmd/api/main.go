package main

import (
	"cma/internal/server"
	"io"
	"log/slog"
	"os"

	sloggin "github.com/samber/slog-gin"
)

func main() {
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "/tmp/cma.log.json"
	}
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	// Delay checking the error for a concise setup
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, f), nil))
	if err != nil {
		logger.Error("could not open log file, logging to stdout only", "err", err)
	}
	logConfig := sloggin.Config{
		WithUserAgent:      false,
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  false,
		WithResponseBody:   true,
		WithResponseHeader: true,
		WithSpanID:         true,
		WithTraceID:        true,
	}

	server := server.NewServer(logger, logConfig)
	logger.Info("starting server", "addr", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("cannot start server", "err", err)
		os.Exit(1)
	}
}
