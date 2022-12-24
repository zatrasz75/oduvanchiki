package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"oduvanchiki/pkg"
)

var router *gin.Engine

func main() {

	router = gin.Default()
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", pkg.Handler)
	router.GET("/form", pkg.HandlerForm)
	err := router.Run(":3737")
	if err != nil {
		return
	}
}
