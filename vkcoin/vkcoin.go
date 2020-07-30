package vkcoin

import "log"

// Merchant - модель платёжного клиента
type Merchant struct {
	id  int
	key string
}

func (merchant Merchant) show() {
	log.Println(merchant.id, merchant.key)
}

func init() {
	log.Println("VKCOIN: DEVELOPED BY BOOST BOTS ( vk.com/boostbots )")
}
