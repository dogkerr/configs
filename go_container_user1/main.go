/*
ini buat tes logging


*/
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, "login bang")
}

type resp struct {
	Message string `json:"message"`
	UserId string `json:"belongs_to_user_id"`
}


func BadRequestHandler(c *gin.Context) {
	c.JSON(http.StatusBadRequest, resp{Message: "Gagal validasi bang", UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"})
}


func UnauthorizedHandler(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, resp{Message: "Lu belum login bang", UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"})
}

func ForbiddenHandler(c *gin.Context) {
	c.JSON(http.StatusForbidden, resp{Message: "Lu gak boleh ke sini bang", UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"})
}

func ServerErrorHandler(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, resp{Message: "Maaf bang developer kita gblg, kode buatannya error semua", UserId: "18d2e020-538d-449a-8e9c-02e4e5cf41111"})
}


func main() {
	r := gin.Default()
	// define the routes
	r.GET("/auth", AuthHandler)
	r.POST("/bad", BadRequestHandler)
	r.GET("/unauthorized", UnauthorizedHandler)
	r.GET("/forbidden", ForbiddenHandler)
	r.GET("/serverError", ServerErrorHandler)

	err := r.Run(":8231")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
