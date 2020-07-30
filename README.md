# VkCoin

VkCoin - пакет для работы с vkc, разрабатываемый для упрощения работы с API VKCOIN.

## Installation / Установка

```bash
go get -u "github.com/golangman/vkcoin/vkcoin"
```

## Usage / Использование

Ниже располагаются примеры использования


### GetBalance / Получение баланса пользователя
```golang
package main

import (
	"log"

	"github.com/golangman/vkcoin/vkcoin"
)

func main() {

	merchant := vkcoin.Merchant{
		ID:  0,              // ID vkcoin аккаунта, с которого был получен токен
		Key: "access_token", // ACCESS_TOKEN аккаунта
	}

	balance, err := merchant.GetBalance(1) // 0 - id для получения баланса текущего аккаунта

	if err != nil {

		log.Fatal(err)

	}

	log.Println(balance)

}
```