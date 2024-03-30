/*
ini buat tes logging


*/
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
//011212999


type response struct {
	Error string `json:"error" example:"message"`
	UserId string `json:"belongs_to_user_id"`

}


func ErrorResponse(c *gin.Context, code int, msg string, userId string) {
	c.AbortWithStatusJSON(code, response{msg, userId})
}




func AuthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, "login bang")
}



func BadRequestHandler(c *gin.Context) {
	ErrorResponse(c, http.StatusBadRequest, "Gagal validasi bang", "18d2e020-538d-449a-8e9c-011212999")
}


func UnauthorizedHandler(c *gin.Context) {
	ErrorResponse(c,http.StatusUnauthorized, "Lu belum login bang",  "18d2e020-538d-449a-8e9c-011212999")
}

func ForbiddenHandler(c *gin.Context) {
	ErrorResponse(c,http.StatusForbidden, "Lu gak boleh ke sini bang",  "18d2e020-538d-449a-8e9c-011212999")

}

func ServerErrorHandler(c *gin.Context) {
	ErrorResponse(c,http.StatusInternalServerError, "Maaf bang developer kita gblg, kode buatannya error semua",  "18d2e020-538d-449a-8e9c-011212999")
}



func main() {
	r := gin.Default()
	// define the routes
	r.GET("/auth", AuthHandler)
	r.POST("/bad", BadRequestHandler)
	r.GET("/unauthorized", UnauthorizedHandler)
	r.GET("/forbidden", ForbiddenHandler)
	r.GET("/serverError", ServerErrorHandler)

	err := r.Run(":8232")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
