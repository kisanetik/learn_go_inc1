package urlmaker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goware/urlx"
	"github.com/kisanetik/learn_go_inc1/config"
	"github.com/kisanetik/learn_go_inc1/internal/storage"
)

const letters = 8

func RandomString() string {
	literals := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	l := make([]rune, letters)
	for i := range l {
		pos := rand.Intn(len(literals)) + int(time.Now().UnixNano()/int64(time.Millisecond)%10)
		if pos > (len(literals) - 1) {
			pos = pos - len(literals)
		}
		if pos < 0 {
			pos = pos * pos
		}
		l[i] = literals[pos]
	}
	return string(l)
}

func CompressURL(url string) string {
	rand := RandomString()
	short := fmt.Sprintf("%s/%s", makeHostFromConfig(), rand)
	record := storage.URLData{UUID: rand, ShortURL: short, OriginalURL: url}
	storage.AddToConfig(record)

	return short
}

func makeHostFromConfig() string {
	_, port := config.LoadConfig()

	url, _ := urlx.Parse(config.GetConf().BaseURL)
	host, port1, _ := urlx.SplitHostPort(url)
	if port1 != "" {
		return fmt.Sprintf("%s://%s:%s", url.Scheme, host, port1)
	}
	return config.GetConf().BaseURL + port
}
