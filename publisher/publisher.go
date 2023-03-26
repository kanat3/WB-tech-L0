package publisher

import (
	"WB-tech-L0/orders"
	"WB-tech-L0/settings"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/stan.go"
)

type PubConfig struct {
	PubName string
	Channel string
	Cluster string
}

type Publisher struct {
	Settings PubConfig
	Data     orders.OrderJSON
}

// Push msg every 15sec

func (pub *Publisher) Run() {
	go func() {
		for {
			pub.Data.Order_uid = randId()
			fmt.Printf("\nPublisher: created data with id: %v\n", pub.Data.Order_uid)
			// generate id for new data
			error := pub.CreateMsg([]byte(`{` + "\n" +
				`"id": "` + fmt.Sprintf("%s", pub.Data.Order_uid) + `",` + "\n" +
				`"order_uid": "b563feb7b2b84b6test",` + "\n" +
				`"track_number": "WBILMTESTTRACK",` + "\n" +
				`"entry": "WBIL",` + "\n" +
				`"delivery": {` + "\n" +
				`"name": "Test Testov",` + "\n" +
				`"phone": "+9720000000",` + "\n" +
				`"zip": "2639809",` + "\n" +
				`"city": "Kiryat Mozkin",` + "\n" +
				`"address": "Ploshad Mira 15",` + "\n" +
				`"region": "Kraiot",` + "\n" +
				`"email": "test@gmail.com"},` + "\n" +
				`"payment": {` + "\n" +
				`"transaction": "b563feb7b2b84b6test",` + "\n" +
				`"request_id": "",` + "\n" +
				`"currency": "USD",` + "\n" +
				`"provider": "wbpay",` + "\n" +
				`"amount": 1817,` + "\n" +
				`"payment_dt": 1637907727,` + "\n" +
				`"bank": "alpha",` + "\n" +
				`"delivery_cost": 1500,` + "\n" +
				`"goods_total": 317,` + "\n" +
				`"custom_fee": 0` + "\n" +
				`},` + "\n" +
				`"items": [` + "\n" +
				`{` + "\n" +
				`"chrt_id": 9934930,` + "\n" +
				`"track_number": "WBILMTESTTRACK",` + "\n" +
				`"price": 453,` + "\n" +
				`"rid": "ab4219087a764ae0btest",` + "\n" +
				`"name": "Mascaras",` + "\n" +
				`"sale": 30,` + "\n" +
				`"size": "0",` + "\n" +
				`"total_price": 317,` + "\n" +
				`"nm_id": 2389212,` + "\n" +
				`"brand": "Vivienne Sabo",` + "\n" +
				`"status": 202` + "\n" +
				`}` + "\n" +
				`],` + "\n" +
				`"locale": "en",` + "\n" +
				`"internal_signature": "",` + "\n" +
				`"customer_id": "test",` + "\n" +
				`"delivery_service": "meest",` + "\n" +
				`"shardkey": "9",` + "\n" +
				`"sm_id": 99,` + "\n" +
				`"date_created": "2021-11-26T06:22:19Z",` + "\n" +
				`"oof_shard": "1"` + "\n" +
				`}` + "\n"))
			if error != nil {
				fmt.Println(error.Error())
			}
			time.Sleep(15 * time.Second)
		}
	}()
	end := make(chan os.Signal, 1)
	signal.Notify(end, os.Interrupt)

	<-end
}

func (pub *Publisher) CreateMsg(msg []byte) error {
	// Connecting
	sc, error := stan.Connect(pub.Settings.Cluster, pub.Settings.PubName, stan.NatsURL("nats1://localhost:4222"))
	// Close connection after publishing
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Publisher: can't connect to the cluster\n")
	}
	defer sc.Close()
	// Create a msg
	error = sc.Publish(pub.Settings.Channel, msg)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Publisher: can't publish msg\n")
	}
	fmt.Printf("'%s' Should see a message in cluster '%s'\n", pub.Settings.Channel, pub.Settings.Cluster)
	return nil
}

func randId() string {
	newId := fmt.Sprintf("%d", rand.Int()%3000000)
	return newId
}

func (pub *Publisher) New() {
	pub.Settings.CreatePublisher()
	pub.Data.New("orders/model.json")
	// General uid
	pub.Data.Order_uid = "1"
}

func (pub *PubConfig) CreatePublisher() {
	//Try read config
	settings.NewConfig(pub, "publisher/publisher.cfg")
}
