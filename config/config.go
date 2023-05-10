package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v8"
)

type cfg struct {
	SERVER_ADDRESS string `env:"SERVER_ADDRESS" envDefault:""`
	BASE_URL       string `env:"BASE_URL" envDefault:""`
}

var (
	HTTPAddr *string
	BaseURL  *string
)

func init() {
	HTTPAddr = flag.String("a", "localhost:8080", "Server address, default is localhost:8080")
	BaseURL = flag.String("b", "http://localhost", "Base URL, default is http://localhost")
	var conf = cfg{
		SERVER_ADDRESS: *HTTPAddr,
		BASE_URL:       *BaseURL,
	}
	if err := env.Parse(&conf); err != nil {
		fmt.Println("failed:", err)
	}
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(*HTTPAddr)

	return host, port
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
