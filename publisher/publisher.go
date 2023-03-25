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
	Data     orders.Order
}

// Push msg and waiting 30 sec

func (pub *Publisher) Run() {
	fmt.Println("Here...")
	go func() {
		for {
			fmt.Println("Ok running publisher...")
			// generate id for new data
			pub.Data.Order_uid = randId()
			pub.CreateMsg([]byte(fmt.Sprintf("%v", pub.Data)))
			fmt.Printf("\npub.Data: %v\n", pub.Data)
			time.Sleep(10 * time.Second)
		}
	}()
	end := make(chan os.Signal, 1)
	signal.Notify(end, os.Interrupt)

	<-end
}

func (pub *Publisher) CreateMsg(msg []byte) error {
	// Connecting
	fmt.Printf("Config: %v\n", pub.Settings)
	sc, error := stan.Connect(pub.Settings.Cluster, pub.Settings.PubName, stan.NatsURL("nats1://localhost:4222"))
	// Close connection after publishing
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't connect to the cluster")
	}
	defer sc.Close()
	// Create a msg
	fmt.Println("Trying to create msg")
	sc.Publish(pub.Settings.Channel, msg)
	fmt.Printf("'%s' Should see a message in cluster '%s'\n", pub.Settings.Channel, pub.Settings.Cluster)
	return nil
}

func randId() string {
	newId := fmt.Sprintf("%d", rand.Int()%3000000)
	return newId
}

func (pub *Publisher) New() {
	fmt.Printf("%s", "Trying to create settings in pub...\n")
	pub.Settings.CreatePublisher()
	pub.Data.New()
	// General uid
	pub.Data.Order_uid = "1"
}

func (pub *PubConfig) CreatePublisher() {
	//Try read config
	settings.NewConfig(pub, "publisher/publisher.cfg")
}
