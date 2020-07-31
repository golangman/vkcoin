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
		ID:  0,              // ID vkcoin аккаунта, с которого был получен access_token
		Key: "access_token", // ACCESS_TOKEN аккаунта ( vkc )
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

### Send / Передача коинов пользователю
```golang
err := merchant.Send(1, 1000) // Перевод 1000 коинов пользователю с id=1

if err != nil {

log.Fatal(err)

}
```

### GetTransactions / Получение списка транзакций
```golang
data, err := merchant.GetTransactions()

if err != nil {

	log.Fatal(err)

}

log.Println(data)

}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Contacts

#### vk.com/boostbots
#### vk.com/golangman
#### vk.com/anetcod
#### golangman@aol.com

## License
[MIT](https://choosealicense.com/licenses/mit/)