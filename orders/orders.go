package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type OrderJSON struct {
	Order_uid string `json:"id"`
	DataJSON  string `json:"data"`
}

type Order struct {
	Order_uid    string `json:"order_uid"`
	Track_number string `json:"track_number"`
	Entry        string `json:"entry"`
	Delivery     struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	}
	Payment struct {
		Transaction   string `json:"transaction"`
		Request_id    string `json:"request_id"`
		Currency      string `json:"currency"`
		Provider      string `json:"provider"`
		Amount        int    `json:"amount"`
		Payment_dt    int    `json:"payment_dt"`
		Bank          string `json:"bank"`
		Delivery_cost int    `json:"delivery_cost"`
		Goods_total   int    `json:"goods_total"`
		Custom_fee    int    `json:"custom_fee"`
	}
	Items []struct {
		Chrt_id      int    `json:"chrt_id"`
		Track_number string `json:"track_number"`
		Price        int    `json:"price"`
		Rid          string `json:"rid"`
		Name         string `json:"name"`
		Sale         int    `json:"sale"`
		Size         string `json:"size"`
		Total_price  int    `json:"total_price"`
		Nm_id        int    `json:"nm_id"`
		Brand        string `json:"brand"`
		Status       int    `json:"status"`
	}
	Locale                   string `json:"locale"`
	Internal_signaturestring string `json:"internal_signature"`
	Customer_id              string `json:"customer_id"`
	Delivery_service         string `json:"delivery_service"`
	Shardkey                 string `json:"shardkey"`
	Sm_id                    int    `json:"sm_id"`
	Date_created             string `json:"date_created"`
	Oof_shard                string `json:"oof_shard"`
}

// for testing nats

func (order *Order) New() error {
	file, error := os.Open("orders/model.json")
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't open the file\n")
	}
	defer file.Close()

	data, error := ioutil.ReadAll(file)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't read headers\n")
	}
	json.Unmarshal(data, &order)
	return nil
}

// for publisher

func (order *OrderJSON) New(fname string) error {
	file, error := os.Open(fname)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't open the file\n")
	}
	defer file.Close()

	data, error := ioutil.ReadAll(file)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't read headers\n")
	}
	order.DataJSON = string(data)
	order.RandomID()
	return nil
}

func (order *OrderJSON) NewById(fname string, id string) error {
	file, error := os.Open(fname)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't open the file\n")
	}
	defer file.Close()

	data, error := ioutil.ReadAll(file)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't read headers\n")
	}
	order.DataJSON = string(data)
	order.Order_uid = id
	return nil
}

func (order *OrderJSON) RandomID() {
	id := fmt.Sprintf("%d", rand.Int()%3000000)
	fmt.Printf("Generated id: %s\n", id)
	order.Order_uid = id
}
