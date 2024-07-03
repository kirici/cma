package server

import (
	"cma/internal/model"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func (s *Server) RegisterRoutes(logger *slog.Logger, logConfig sloggin.Config) http.Handler {
	r := gin.New()
	r.Use(sloggin.NewWithConfig(logger, logConfig))
	r.Use(gin.Recovery())

	r.GET("/", s.HelloWorldHandler)
	r.GET("/healthz", s.healthHandler)
	r.GET("/order/:id", s.orderHandler)

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

func (s *Server) orderHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := model.Order{
		Id: id,
	}
	c.JSON(http.StatusOK, s.db.AddOrder(order))
}
