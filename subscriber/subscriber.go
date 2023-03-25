package subscriber

import (
	"WB-tech-L0/datacache"
	"WB-tech-L0/orders"
	"WB-tech-L0/settings"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
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

func (s *SubSettings) New() {
	settings.NewConfig(s, "subscriber/subscriber.cfg")
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
		return errors.New("Can't connect to database")
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
		return errors.New("Can't push to database")
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
		return errors.New("Can't get from database")
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
