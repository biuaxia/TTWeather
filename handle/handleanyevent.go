package main

import (
	"TTWeather/common"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

// main, handle all types of events.
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

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		fmt.Printf("OnText: %+v\n", m)
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
		fmt.Printf("OnPhoto: %+v\n", m)
	})

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		// channel posts only
		fmt.Printf("OnChannelPost: %+v\n", m)
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		// incoming inline queries
		fmt.Printf("OnQuery: %+v\n", q)
	})

	b.Start()
}
