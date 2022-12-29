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
	router.GET("/form", pkg.GetHandlerForm)
	router.POST("/form", pkg.PostHandlerForm)
	err := router.Run(":3737")
	if err != nil {
		return
	}

}

// https://www.youtube.com/watch?v=BqhjWav0MuA 7 min
