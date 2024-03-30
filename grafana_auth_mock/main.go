package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func TesHandler(c *gin.Context) {
	fmt.Println("asdosajodsaasoasasdd")
	
	c.JSON(http.StatusOK, "asdsasdsd")
 }
 
 
 

func main() {
	r := gin.Default()
	// define the routes
	r.GET("/auth", TesHandler)
	err := r.Run(":8083")
	if err != nil {
	   log.Fatalf("error: %s", err)
	}
 }
 