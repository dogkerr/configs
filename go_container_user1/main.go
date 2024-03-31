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
	"time"

	"tes/httpserver"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error  string `json:"error" example:"message"`
	UserId string `json:"belongs_to_user_id"`
}

type succsessResponse struct {
	Message string `json:"message"`
	UserId  string `json:"belongs_to_user_id"`
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
	start := time.Now()
	time.Sleep(time.Millisecond * 200) // simulasi operasi gajelas
	end := time.Now()
	latency := end.Sub(start)
	su.l.Info("API request to /api/v1/auth completed successfully- 18d2e020-538d-449a-8e9c-02e4e5cf41111",
		logger.LogMetadata{Clientid: c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: 200, BodySize: c.Writer.Size(),
			Path:    c.Request.URL.Path,
			Latency: latency.String()})
	resp := succsessResponse{Message: "API request to /api/v1/auth completed successfully",
		UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"}
	c.JSON(http.StatusOK, resp)
}

// c.JSON(http.StatusOK, "API request to /api/v1/auth completed successfully")

func (su *cok) ErrorHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Millisecond * 200) // simulasi operasi gajelas
	end := time.Now()
	latency := end.Sub(start)
	su.l.Warn("gagal validasi - 18d2e020-538d-449a-8e9c-02e4e5cf41111",
		logger.LogMetadata{Clientid: c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: 400, BodySize: c.Writer.Size(),
			Path:    c.Request.URL.Path,
			Latency: latency.String()})
	ErrorResponse(c, http.StatusBadRequest, "gagal validasi ", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) WarnHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Millisecond * 200) // simulasi operasi gajelas
	end := time.Now()
	latency := end.Sub(start)
	su.l.Warn("CPU usage warning - 18d2e020-538d-449a-8e9c-02e4e5cf41111",
		logger.LogMetadata{Clientid: c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: 401, BodySize: c.Writer.Size(),
			Path:    c.Request.URL.Path,
			Latency: latency.String()})

	ErrorResponse(c, http.StatusUnauthorized, "CPU usage warning", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) FatalHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Millisecond * 200) // simulasi operasi gajelas
	end := time.Now()
	latency := end.Sub(start)
	su.l.Error(errors.New("no space available for write operations - 18d2e020-538d-449a-8e9c-02e4e5cf41111"),
		logger.LogMetadata{Clientid: c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: 500, BodySize: c.Writer.Size(),
			Path:    c.Request.URL.Path,
			Latency: latency.String()})
	ErrorResponse(c, http.StatusInternalServerError, "no space available for write operations", "18d2e020-538d-449a-8e9c-02e4e5cf41111")

}

func (su *cok) DebugHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Millisecond * 200) // simulasi operasi gajelas
	end := time.Now()
	latency := end.Sub(start)
	su.l.Debug("created order ps5 - 18d2e020-538d-449a-8e9c-02e4e5cf41111",
		logger.LogMetadata{Clientid: c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: 201, BodySize: c.Writer.Size(),
			Path:    c.Request.URL.Path,
			Latency: latency.String()})
	resp := succsessResponse{Message: "created order ps5", UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"}
	c.JSON(http.StatusCreated, resp)
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
