package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v8"
)

type cfg struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:""`
	BaseURL       string `env:"BASE_URL" envDefault:""`
}

var conf cfg

func init() {
	HTTPAddr := flag.String("a", "localhost:8080", "Server address, default is localhost:8080")
	BaseURL := flag.String("b", "http://localhost", "Base URL, default is http://localhost")
	conf = cfg{
		ServerAddress: *HTTPAddr,
		BaseURL:       *BaseURL,
	}
	if err := env.Parse(&conf); err != nil {
		fmt.Println("failed:", err)
	}
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(conf.ServerAddress)

	return host, port
}

func GetConf() cfg {
	return conf
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
