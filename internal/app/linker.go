package linker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goware/urlx"
	"github.com/kisanetik/learn_go_inc1/config"
)

func CompressURL(url string) string {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(tFile.Name(), []byte(url), 0644)

	return string(makeHostFromConfig() + "/" + filepath.Base(tFile.Name()))
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
