# VkCoin

VkCoin - пакет для работы с vkc, разрабатываемый для упрощения работы с API VKCOIN.

## Installation / Установка

```bash
go get -u "github.com/golangman/vkcoin/vkcoin"
```

## Usage / Использование

Ниже располагаются примеры использования


### Init Merchant / Инициализация платёжного сервиса
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

}
```

### GetBalance / Получение баланса пользователя
```golang
balance, err := merchant.GetBalance(1) // 0 - id для получения баланса текущего аккаунта

if err != nil {

log.Fatal(err)

}

log.Println(balance)
```

### GetBalance / Передача коинов пользователю
```golang
err := merchant.Send(1, 1000) // Перевод 1000 коинов пользователю с id=1

if err != nil {

log.Fatal(err)

}
```