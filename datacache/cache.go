package datacache

import (
	"WB-tech-L0/orders"
	"WB-tech-L0/settings"
	"errors"
	"sync"
)

type DataCache struct {
	Data map[string]orders.OrderJSON
	sync.RWMutex
}

type DbSettings struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (s *DbSettings) New() {
	//Try read config
	settings.NewConfig(s, "datacache/database.cfg")
}

func New() *DataCache {
	bd := &DataCache{}
	bd.Data = make(map[string]orders.OrderJSON)
	return bd
}

func (bd *DataCache) Insert(order *orders.OrderJSON) {
	bd.Lock()
	defer bd.Unlock()
	bd.Data[order.Order_uid] = *order
}

func (bd *DataCache) InsertByLabel(id string, json string) {
	var order orders.OrderJSON
	order.Order_uid = id
	order.DataJSON = json
	if id == "" || json == "" {
		return
	}
	bd.Lock()
	defer bd.Unlock()
	bd.Data[id] = order
}

func (bd *DataCache) Get(key string) (orders.OrderJSON, error) {
	bd.RLock()
	defer bd.RUnlock()
	if bd.Data[key].Order_uid == "" {
		return bd.Data[key], errors.New("Here no such order with your ID :(")
	}
	return bd.Data[key], nil
}
