package server

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func (s *Server) RegisterRoutes(logger *slog.Logger, logConfig sloggin.Config) http.Handler {
	r := gin.New()
	r.Use(sloggin.NewWithConfig(logger, logConfig))
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
