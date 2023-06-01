package urlmaker

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/goware/urlx"
	"github.com/kisanetik/learn_go_inc1/config"
	"go.uber.org/zap"
)

const letters = 8

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Mem struct {
	memory map[string]URLData
}

var cache *Mem

func RandomString() string {
	literals := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	l := make([]rune, letters)
	for i := range l {
		l[i] = literals[rand.Intn(len(literals))]
	}

	return string(l)
}

func getCache() (*Mem, error) {
	if nil == cache {
		file, err := os.OpenFile(config.GetConf().FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
	}

	return cache
}

func CompressURL(url string) string {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		logger, err := zap.NewDevelopment()
		if err != nil {
			// вызываем панику, если ошибка
			panic("cannot initialize zap")
		}
		defer logger.Sync()
		sugar := logger.Sugar()
		sugar.Errorf("Error when create temporary file: %s", err)
	}
	os.WriteFile(tFile.Name(), []byte(url), 0644)

	return fmt.Sprintf("%s/%s", makeHostFromConfig(), filepath.Base(tFile.Name()))
}

func makeHostFromConfig() string {
	// var link strings.Builder
	_, port := config.LoadConfig()

	url, _ := urlx.Parse(config.GetConf().BaseURL)
	host, port1, _ := urlx.SplitHostPort(url)
	if port1 != "" {
		return fmt.Sprintf("%s://%s:%s", url.Scheme, host, port1)
	}
	return config.GetConf().BaseURL + port
}
