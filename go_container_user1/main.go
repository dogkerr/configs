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

func (su *cok) AuthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, "login bang")
}

func (su *cok) BadRequestHandler(c *gin.Context) {
	su.l.Error(errors.New("gagal validasi bang - 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - BadRequestHandler")
	ErrorResponse(c, http.StatusBadRequest, "Gagal validasi bang", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) UnauthorizedHandler(c *gin.Context) {
	su.l.Error(errors.New("lu belum login bang - 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - UnauthorizedHandler")

	ErrorResponse(c, http.StatusUnauthorized, "Lu belum login bang", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func (su *cok) ForbiddenHandler(c *gin.Context) {
	su.l.Error(errors.New("lu gak boleh ke sini bang - 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - ForbiddenHandler")
	ErrorResponse(c, http.StatusForbidden, "Lu gak boleh ke sini bang", "18d2e020-538d-449a-8e9c-02e4e5cf41111")

}

func (su *cok) ServerErrorHandler(c *gin.Context) {
	su.l.Error(errors.New("maaf bang developer kita gblg, kode buatannya error semua - 18d2e020-538d-449a-8e9c-02e4e5cf41111"), "http - v1 - ServerErrorHandler")
	ErrorResponse(c, http.StatusInternalServerError, "Maaf bang developer kita gblg, kode buatannya error semua", "18d2e020-538d-449a-8e9c-02e4e5cf41111")
}

func main() {
	l := logger.New("debug")
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	cok := newCok(l)

	// define the routes
	r.GET("/auth", cok.AuthHandler)
	r.POST("/bad", cok.BadRequestHandler)
	r.GET("/unauthorized", cok.UnauthorizedHandler)
	r.GET("/forbidden", cok.ForbiddenHandler)
	r.GET("/serverError", cok.ServerErrorHandler)

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
