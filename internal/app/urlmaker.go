package urlmaker

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goware/urlx"
	"github.com/kisanetik/learn_go_inc1/config"
	"go.uber.org/zap"
)

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
