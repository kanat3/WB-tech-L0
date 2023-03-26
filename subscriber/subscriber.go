package subscriber

import (
	"WB-tech-L0/datacache"
	"WB-tech-L0/orders"
	"WB-tech-L0/settings"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

type SubSettings struct {
	Name      string
	Cluster   string
	Channel   string
	CacheFile string
}

type Subscriber struct {
	SubSettings SubSettings
	Cache       *datacache.DataCache
	DbSettings  datacache.DbSettings
}

func (sub *Subscriber) Run() error {
	sc, error := stan.Connect(sub.SubSettings.Cluster, sub.SubSettings.Name, stan.NatsURL("nats1://localhost:4222"))
	// Close connection after publishing
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: can't connect to the cluster\n")
	}

	db, error := sub.DbConnect()
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: can't connect to the database\n")
	}

	defer db.Close()

	// Restore database into cache
	error = sub.DbToCache(db)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: database didn't restore\n")
	}

	// Parse and append correct messages to cache
	res, error := sc.Subscribe(sub.SubSettings.Channel, func(msg *stan.Msg) {
		order := orders.OrderJSON{}
		error = json.Unmarshal(msg.Data, &order)
		if error != nil {
			fmt.Println(error.Error())
		}
		order.DataJSON = string(msg.Data)
		fmt.Printf("Subscriber: get a message with id: %s\n", order.Order_uid)
		if error == nil {
			// Append message to cache
			sub.Cache.Insert(&order)
			// Push order to db
			error = sub.PushOrder(order, db)
			if error != nil {
				fmt.Println(error.Error())
			}
		}
	})
	fmt.Printf("Connected to clusterID: [%s] clientID: [%s]\n", sub.SubSettings.Cluster, sub.SubSettings.Name)

	// Unsubscribe if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Ctrl+C... closing connection...")
	res.Close()
	sc.Close()

	return nil

}

func (sub *SubSettings) New() {
	settings.NewConfig(sub, "subscriber/subscriber.cfg")
}

func New() *Subscriber {
	sub := &Subscriber{}
	sub.SubSettings.New()
	sub.DbSettings.New()

	// need to create map
	sub.Cache = datacache.New()
	return sub
}

func (sub *Subscriber) AddDbToCache(db *datacache.DataCache) {
	sub.Cache = db
}

func (sub *Subscriber) PingDb() error {
	ping := "host=" + sub.DbSettings.Host + " port=" + sub.DbSettings.Port + " user=" + sub.DbSettings.User +
		" password=" + sub.DbSettings.Password + " sslmode=disable"
	db, error := sql.Open("postgres", ping)

	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: can't connect to database")
	}

	error = db.Ping()
	return nil
}

func (sub *Subscriber) DbConnect() (*sql.DB, error) {
	ping := "host=" + sub.DbSettings.Host + " port=" + sub.DbSettings.Port + " user=" + sub.DbSettings.User +
		" password=" + sub.DbSettings.Password + " sslmode=disable"

	db, error := sql.Open("postgres", ping)

	if error != nil {
		fmt.Println(error.Error())
		return nil, error
	}
	error = db.Ping()

	if error != nil {
		fmt.Println(error.Error())
		return nil, error
	}
	return db, error
}

func (sub *Subscriber) PushOrder(order orders.OrderJSON, db *sql.DB) error {
	db, error := sub.DbConnect()

	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: can't push to database")
	}

	_, error = db.Exec("call push('" + order.Order_uid + "', '" + order.DataJSON + "');")

	if error != nil {
		fmt.Println(error.Error())
		return error
	}
	return nil
}

func (sub *Subscriber) DbToCache(db *sql.DB) error {
	error := sub.PingDb()

	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Subscriber: can't get from database")
	}

	data, error := db.Query("select * from orders;")

	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	defer data.Close()

	for data.Next() {
		order := orders.OrderJSON{}
		error = data.Scan(&order.Order_uid, &order.DataJSON)
		id := order.Order_uid
		json := order.DataJSON

		if error != nil {
			fmt.Println(error.Error())
			return error
		}

		sub.Cache.InsertByLabel(id, json)
	}
	return nil
}
