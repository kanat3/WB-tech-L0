package main

import (
	//"Work/WB-tech-L0/datacache"

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

	// pull every note from database to subscriber.cache
	error = sub.DbToCache(db)

	if error != nil {
		fmt.Printf(error.Error())
		return
	}

	go sub.Run()

	// create data to push
	/*
		var order orders.OrderJSON

		order.New("orders/model.json")

		sub.PushOrder(order, db)
		fmt.Printf("Data id generated: \n %v \n", order.Order_uid)
	*/
	// here using get from cache. Ok
	/*
		var order_get orders.OrderJSON

		order_get, error = sub.Cache.GetById(order.Order_uid)

		if error != nil {
			fmt.Printf(error.Error())
			return
		}

		fmt.Printf("Data by getted by subcriber: \n %v", order_get)
	*/
	// is working
	//fmt.Printf("Sublisher all data: \n %v \n %v \n %v \n", sub.SubSettings, sub.DbSettings, sub.Cache)

	// try to push it in database

	//data := datacache.New()

	router = gin.Default()

	router.LoadHTMLFiles("static/index.html", "static/bye_page.html")
	router.Static("/images", "./images")
	router.Static("static/css", "./static/css")

	router.GET("/bye_page", bye)

	var pub publisher.Publisher
	id := ""
	pub.New()

	go pub.Run()

	router.GET("/", func(c *gin.Context) {
		index(c)
		id = getId(c)
		_, error := sub.Cache.GetById(id)
		if error != nil {
			fmt.Print(error.Error())
		} else {
			fmt.Printf("OK: found order with id: %v\n", id)
		}
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

func bye(c *gin.Context) {
	c.HTML(200, "bye_page.html", nil)
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
