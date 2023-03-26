package main

import (
	//"Work/WB-tech-L0/datacache"

	"WB-tech-L0/orders"
	"WB-tech-L0/publisher"
	"WB-tech-L0/subscriber"
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	var sub *subscriber.Subscriber
	// intitialization
	sub = subscriber.New()

	// connect to bd
	db, error := sub.DbConnect()
	if error != nil {
		fmt.Printf(error.Error())
		return
	}

	// get all database to cache
	error = sub.DbToCache(db)

	if error != nil {
		fmt.Printf(error.Error())
		return
	}
	go sub.Run()

	router = gin.Default()

	router.LoadHTMLFiles("static/index.html", "static/bye_page.html")
	router.Static("/images", "./images")
	router.Static("static/css", "./static/css")

	var pub publisher.Publisher
	pub.New()

	go pub.Run()

	var o orders.OrderJSON
	router.GET("/", func(c *gin.Context) {
		index(c)
		id := getId(c)
		o, error = sub.Cache.GetById(id)
		if error != nil {
			fmt.Print(error.Error())
		} else {
			fmt.Printf("OK: found order with id: %v\n", id)
		}
	})

	router.GET("/bye_page", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"id":   o.Order_uid,
			"data": o.DataJSON,
		})
	})

	router.GET("/images", images)
	router.GET("/static/css", page)

	router.Run("127.0.0.1:8090")

	end := make(chan os.Signal, 1)
	signal.Notify(end, os.Interrupt)

	<-end

}

func getId(c *gin.Context) (id string) {
	id, ok := c.GetQuery("data")
	if !ok {
		fmt.Println("Can't get data from form")
		return ""
	}
	return id
}

func images(c *gin.Context) {
	c.HTML(200, "images.html", nil)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func page(c *gin.Context) {
	c.HTML(200, "page.css", nil)
}

func image(c *gin.Context) {
	c.HTML(200, "town.jpg", nil)
}
