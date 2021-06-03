package main

import (
	"TTWeather/common"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

// main, use LongPoller receive user send content, but received '/hello' command,
// will send to user 'Hello World!'.
func main() {
	// Proxy config
	host := "localhost"
	port := 8889

	// Token, form BotFather
	token := "1755685454:AAHEf3BRaYKj-l8oC_m6V8JL02Gwk77mOM4"

	b, err := tb.NewBot(tb.Settings{
		Client: common.BuildClient(host, port),

		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		_, err := b.Send(m.Sender, "Hello World!")
		if err != nil {
			return
		}
	})

	b.Start()
}
