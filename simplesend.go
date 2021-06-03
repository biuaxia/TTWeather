package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"net/url"
	"time"
)

// buildClient is a method used to construct a request proxy.
func buildClient(host string, port int) *http.Client {
	proxyUrl := fmt.Sprintf("http://%s:%d", host, port)

	uri, err := url.Parse(proxyUrl)

	if err != nil {
		log.Fatal("error parsing proxy url: ", err)
	}

	return &http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
}

// main, use LongPoller receive user send content, but received '/hello' command,
// will send to user 'Hello World!'.
func main() {
	// Proxy config
	host := "localhost"
	port := 8889

	// Token, form BotFather
	token := ""

	b, err := tb.NewBot(tb.Settings{
		Client: buildClient(host, port),

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
