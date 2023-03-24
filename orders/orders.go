package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

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

func (o *Order) ReadOrder() error {
	file, error := os.Open("orders/model.json")
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't open the file")
	}
	defer file.Close()

	data, error := ioutil.ReadAll(file)
	if error != nil {
		fmt.Println(error.Error())
		return errors.New("Can't read headers")
	}
	json.Unmarshal(data, &o)
	return nil
}
