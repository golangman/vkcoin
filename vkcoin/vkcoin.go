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

// VkcURL - api domain
const VkcURL = "https://coin-without-bugs.vkforms.ru/merchant/"

// Merchant - модель платёжного клиента
type Merchant struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// SendResponse - модель ответа vkcoin
type SendResponse struct {
	ID      int
	Amount  int
	Current int
}

// GetTransactionsResponse - ...
type GetTransactionsResponse struct {
	Transactions []Transaction `json:"response"`
}

// GetBalanceResponse - ...
type GetBalanceResponse struct {
	Transactions []Transaction `json:"response"`
}

// RequestError - модель ответа vkcoin
type RequestError struct {
	Error VkcError `json:"error"`
}

// Transaction - модель транзакции
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

// VkcError - vkc request error
type VkcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetBalance - получение баланса пользователя по id
func (merchant Merchant) GetBalance(userID int) (int, error) {
	url := VkcURL + "score/"

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

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

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
func (merchant Merchant) Send(toID, amount int) (response *SendResponse, errorVkc *RequestError, err error) {
	url := VkcURL + "send/"

	data, err := jsoniter.Marshal(map[interface{}]interface{}{

		"merchantId": merchant.ID,
		"key":        merchant.Key,
		"toId":       toID,
		"amount":     amount * 1000,
	})

	if err != nil {

		return nil, nil, err

	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {

		return nil, nil, err

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return nil, nil, err

	}

	if err := json.Unmarshal(body, &response); err != nil {

		return nil, nil, err

	}

	if isEmpty(response) {

		err = json.Unmarshal(data, &errorVkc)
		if err != nil {
			return nil, nil, err
		}
		return nil, errorVkc, err
	}

	return nil, nil, nil
}

// GetTransactions - получение списка входящих переводов
func (merchant Merchant) GetTransactions() (response *GetTransactionsResponse, errorVkc *RequestError, err error) {
	url := VkcURL + "tx/"

	data, err := jsoniter.Marshal(map[interface{}]interface{}{

		"merchantId": merchant.ID,
		"key":        merchant.Key,
		"tx":         [1]int{1},
	})

	if err != nil {

		return nil, nil, err

	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, nil, err
	}

	if isEmpty(response) {

		err = json.Unmarshal(data, &errorVkc)
		if err != nil {
			return nil, nil, err
		}
		return nil, errorVkc, err
	}

	return response, nil, nil
}

func init() {
	log.Println("VKCOIN: DEVELOPED BY BOOST BOTS ( vk.com/boostbots )")
}
