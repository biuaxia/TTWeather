package common

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// BuildClient is a method used to construct a request proxy.
func BuildClient(host string, port int) *http.Client {
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
