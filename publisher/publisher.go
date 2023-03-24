package publisher

import (
	"Work/WB-tech-L0/orders"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"

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

func New() *Publisher {
	pub := &Publisher{}
	pub.Settings.CreatePublisher()
	return pub
}

func (pub *Publisher) SimulatePub() error {
	var fakeOrder orders.Order
	error := fakeOrder.ReadOrder()
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't create an order")
	}
	fakeOrder.Order_uid = randId()
	pub.Data = fakeOrder
	return nil
}

func (pub *Publisher) CreateMsg(msg []byte) {
	// Connecting
	sc, error := stan.Connect(pub.Settings.Cluster, pub.Settings.PubName)
	// Close connection after publishing
	defer sc.Close()
	if error != nil {
		fmt.Println("Can't connect to the cluster")
		return
	}
	// Create a msg
	sc.Publish(pub.Settings.Channel, msg)
	fmt.Printf("'%s' Should see a message in cluster '%s'\n", pub.Settings.Channel, pub.Settings.Cluster)
}

func randId() string {
	newId := fmt.Sprintf("%d", rand.Int()%50)
	return newId
}

func (pub *Publisher) New() {
	pub.Settings.CreatePublisher()
}

func (pub *PubConfig) CreatePublisher() {
	//try read config
	file, err := os.Open("publisher/publisher.cfg")
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't open the configuration")
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		panic("Cant't read the structure of the configuration")
	}

	readByte := make([]byte, stat.Size())

	_, err = file.Read(readByte)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't read the configuration")
	}
	err = json.Unmarshal(readByte, &pub)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't read configuratiion data")
	}
}
