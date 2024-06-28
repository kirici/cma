package server

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "/tmp/cma.log.json"
	}
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		slog.Error("could not open log file, logging to stdout only", "err", err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, f), nil))
	config := sloggin.Config{
		WithUserAgent:      false,
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  false,
		WithResponseBody:   true,
		WithResponseHeader: true,
		WithSpanID:         true,
		WithTraceID:        true,
	}

	r := gin.New()
	r.Use(sloggin.NewWithConfig(logger, config))
	r.Use(gin.Recovery())

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
