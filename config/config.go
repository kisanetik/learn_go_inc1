package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v8"
)

type cfg struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

var conf cfg

func init() {
	conf = cfg{
		ServerAddress: "localhost:8080",
		BaseURL:       "http://localhost",
	}
	if err := env.Parse(&conf); err != nil {
		fmt.Println("failed:", err)
	}
	conf.loadFlags()
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(conf.ServerAddress)

	return host, port
}

func (cfg *cfg) loadFlags() {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address, default is localhost:8080")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL, default is http://localhost")
}

func GetConf() cfg {
	return conf
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
