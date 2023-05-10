package linker

import (
	"os"
	"path/filepath"
	"strings"

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

	if strings.Contains(*config.BaseURL, ":") {
		return *config.BaseURL
	}
	return *config.BaseURL + port
}
