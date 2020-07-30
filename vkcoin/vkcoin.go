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
	ID  int
	Key string
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

func init() {
	log.Println("VKCOIN: DEVELOPED BY BOOST BOTS ( vk.com/boostbots )")
}
