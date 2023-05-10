package linker

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goware/urlx"
	"github.com/kisanetik/learn_go_inc1/config"
)

func CompressURL(url string) string {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	os.WriteFile(tFile.Name(), []byte(url), 0644)

	return string(makeHostFromConfig() + "/" + filepath.Base(tFile.Name()))
}

func makeHostFromConfig() string {
	_, port := config.LoadConfig()

	url, _ := urlx.Parse(*config.BaseURL)
	host, port1, _ := urlx.SplitHostPort(url)
	if port1 != "" {
		fmt.Println("Contains " + *config.BaseURL + "/")
		return url.Scheme + "://" + host + ":" + port1
	}
	fmt.Println("NOT contains" + *config.BaseURL)
	return *config.BaseURL + port
}
