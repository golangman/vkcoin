package vkcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

// Merchant - модель платёжного клиента
type Merchant struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// Answer - модель ответа vkcoin
type Answer struct {
	Response interface{} `json:"response"`
}

// Transaction - модель транзакциb
type Transaction struct {
	Amount     string `json:"amount"`
	CreatedAt  int    `json:"created_at"`
	ExternalID int    `json:"external_id"`
	FromID     int    `json:"from_id"`
	ID         int    `json:"id"`
	Payload    int    `json:"payload"`
	ToID       int    `json:"to_id"`
	Type       int    `json:"type"`
}

// GetBalance - получение баланса пользователя по id
func (merchant Merchant) GetBalance(userID int) (int, error) {
	_url := "https://coin-without-bugs.vkforms.ru/merchant/score/"

	if userID == 0 {
		userID = merchant.ID
	}

	data, err := jsoniter.Marshal(map[interface{}]interface{}{

		"merchantId": merchant.ID,
		"key":        merchant.Key,
		"userIds":    [1]int{userID},
	})

	if err != nil {

		return 0, err

	}

	resp, err := http.Post(_url, "application/json", bytes.NewBuffer(data))

	if err != nil {

		return 0, err

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return 0, err

	}

	m := map[string]interface{}{}

	if err := json.Unmarshal(body, &m); err != nil {

		return 0, err

	}

	balanceStr := fmt.Sprintf("%.0f", m["response"].(map[string]interface{})[fmt.Sprintf("%d", userID)])

	balance, err := strconv.Atoi(balanceStr)

	if err != nil {

		return 0, err

	}

	return balance / 1000, nil

}

// Send - передача коинов пользователю
func (merchant Merchant) Send(toID, amount int) error {
	_url := "https://coin-without-bugs.vkforms.ru/merchant/send/"

	data, err := jsoniter.Marshal(map[interface{}]interface{}{

		"merchantId": merchant.ID,
		"key":        merchant.Key,
		"toId":       toID,
		"amount":     amount * 1000,
	})

	if err != nil {

		return err

	}

	resp, err := http.Post(_url, "application/json", bytes.NewBuffer(data))

	if err != nil {

		return err

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return err

	}

	m := map[string]interface{}{}

	if err := json.Unmarshal(body, &m); err != nil {

		return err

	}

	if _, ok := m["error"]; ok {

		return err

	}

	return nil
}

// GetTransactions - получение списка входящих переводов
func (merchant Merchant) GetTransactions() ([]Transaction, error) {
	_url := "https://coin-without-bugs.vkforms.ru/merchant/tx/"

	data, err := jsoniter.Marshal(map[interface{}]interface{}{

		"merchantId": merchant.ID,
		"key":        merchant.Key,
		"tx":         [1]int{1},
	})

	if err != nil {

		return nil, err

	}

	resp, err := http.Post(_url, "application/json", bytes.NewBuffer(data))

	if err != nil {

		return nil, err

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return nil, err

	}

	m := map[string]interface{}{}

	mResponse := []Transaction{}

	if err := json.Unmarshal(body, &m); err != nil {

		return nil, err

	}

	mResponseJSON, err := json.Marshal(m["response"])

	if err != nil {

		return nil, err

	}

	err = json.Unmarshal(mResponseJSON, &mResponse)

	if err != nil {

		return nil, err

	}

	return mResponse, nil

}

func init() {
	log.Println("VKCOIN: DEVELOPED BY BOOST BOTS ( vk.com/boostbots )")
}
