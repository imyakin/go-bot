package telegram

import "github.com/imyakin/go-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

//func New(client *telegram.Client, storage) {
//
//}
