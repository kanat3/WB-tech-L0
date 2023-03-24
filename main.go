package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	//data := datacache.New()
	//pub := publisher.New()

	//go pub.Run()

	fmt.Println(cfg.ServerHost + ":" + cfg.ServerPort)

	router = gin.Default()

	router.LoadHTMLFiles(cfg.HTML + "index.html")
	router.Static("/images", "./images")

	router.GET("/", index)
	router.GET("/images", index)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)

	end := make(chan os.Signal, 1)
	signal.Notify(end, os.Interrupt)

	<-end
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
