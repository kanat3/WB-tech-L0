package main

import (
	//"Work/WB-tech-L0/datacache"

	"WB-tech-L0/subscriber"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	var sub *subscriber.Subscriber

	sub = subscriber.New()

	db, error := sub.DbConnect()

	if error != nil {
		fmt.Printf(error.Error())
		return
	}

	// create data to push

	//var order orders.OrderJSON

	//order.New("orders/model.json")

	//sub.PushOrder(order, db)
	//fmt.Printf("Data id generated: \n %v \n", order.Order_uid)
	// get all database to cache

	sub.DbToCache(db)

	fmt.Printf("Sublisher all data: \n %v \n %v \n %v \n", sub.SubSettings, sub.DbSettings, sub.Cache)

	// try to push it in database

	//data := datacache.New()
	/*
		var pub publisher.Publisher
		pub.New()

		pub.Run()

		var cfg settings.Settings
		settings.NewConfig(cfg, "settings.cfg")
		fmt.Println(cfg.ServerHost + ":" + cfg.ServerPort)
	*/
	/*
		router = gin.Default()

		router.LoadHTMLFiles(cfg.Static+"index.html", cfg.Static+"bye_page.html")
		router.Static("/images", cfg.Images)
		router.Static("static/css", "./static/css")

		router.GET("/bye_page", bye)
		router.GET("/", func(c *gin.Context) {
			index(c)
			pub.Data.Order_uid = getId(c)
			fmt.Println("id is: ", pub.Data.Order_uid)
		})

		router.GET("/images", images)
		router.GET("/static/css", page)

		router.Run(cfg.ServerHost + ":" + cfg.ServerPort)

		end := make(chan os.Signal, 1)
		signal.Notify(end, os.Interrupt)

		<-end
	*/
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
