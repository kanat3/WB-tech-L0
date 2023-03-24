package datacache

import (
	"Work/WB-tech-L0/orders"
	"errors"
	"sync"
)

type DataCache struct {
	Data map[string]orders.Order
	sync.RWMutex
}

func New() *DataCache {
	bd := &DataCache{}
	bd.Data = make(map[string]orders.Order)
	return bd
}

func (bd *DataCache) Insert(o *orders.Order) {
	bd.Lock()
	defer bd.Unlock()
	bd.Data[o.Order_uid] = *o
}

func (bd *DataCache) Get(key string) (orders.Order, error) {
	bd.RLock()
	defer bd.RUnlock()
	if bd.Data[key].Order_uid == "" {
		return bd.Data[key], errors.New("Here no such order with your ID :(")
	}
	return bd.Data[key], nil
}
