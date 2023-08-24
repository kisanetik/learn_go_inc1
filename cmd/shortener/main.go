package main

import (
	"github.com/kisanetik/learn_go_inc1/internal/config"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Can't read config: %w", err)
	}
}
