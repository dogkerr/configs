/*
ini buat tes logging
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tes/logger"

	"tes/httpserver"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error  string `json:"error" example:"message"`
	UserId string `json:"belongs_to_user_id"`
}

func ErrorResponse(c *gin.Context, code int, msg string, userId string) {
	c.AbortWithStatusJSON(code, response{msg, userId})
}

type cok struct {
	l logger.Interface
}

func newCok(l logger.Interface) *cok {
	return &cok{l}
}

func (su *cok) InfoHandler(c *gin.Context) {
	su.l.Info("API request to /api/v1/auth completed successfully- 18d2e020-538d-449a-8e9c-02e4e5cf41111", "http - v1 - ErrorHandler")

	c.JSON(http.StatusOK, "API request to /api/v1/auth completed successfully")
}

func (su *cok) ErrorHandler(c *gin.Context) {
	su.l.Error(errors.New("connection timeout- 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - ErrorHandler")
	ErrorResponse(c, http.StatusBadRequest, "connection timeout", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) WarnHandler(c *gin.Context) {
	su.l.Warn("CPU usage warning - 18d2e020-538d-449a-8e9c-02e4e5cf41111", "http - v1 - WarnHandler")

	ErrorResponse(c, http.StatusUnauthorized, "CPU usage warning", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) FatalHandler(c *gin.Context) {
	// su.l.Fatal(errors.New("no space available for write operations - 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - FatalHandler")
	ErrorResponse(c, http.StatusForbidden, "no space available for write operations", "18d2e020-538d-449a-8e9c-02e4e5cf41111")

}

func (su *cok) DebugHandler(c *gin.Context) {
	su.l.Debug("select user1 from the database - 18d2e020-538d-449a-8e9c-02e4e5cf41111", "http - v1 - DebugHandler")
	ErrorResponse(c, http.StatusInternalServerError, "select user1 from the database", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func main() {
	l := logger.New()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	cok := newCok(l)

	// define the routes
	r.GET("/info", cok.InfoHandler)
	r.POST("/error", cok.ErrorHandler)
	r.GET("/warn", cok.WarnHandler)
	r.GET("/fatal", cok.FatalHandler)
	r.GET("/debug", cok.DebugHandler)

	httpServer := httpserver.New(r, httpserver.Port("8231"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
