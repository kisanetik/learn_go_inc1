package config

import (
	"flag"
	"strings"
)

var (
	HTTPAddr *string
	BaseURL  *string
)

func init() {
	HTTPAddr = flag.String("a", "localhost:8080", "Server address, default is localhost:8080")
	BaseURL = flag.String("b", "http://localhost", "Base URL, default is http://localhost")
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(*HTTPAddr)

	return host, port
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
