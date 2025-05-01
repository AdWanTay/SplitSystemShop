package main

import (
	"SplitSystemShop/internal/app"
	"SplitSystemShop/internal/config"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("get config: %v", err)
	}

	err = app.App(cfg)
	if err != nil {
		log.Fatalf("load app: %v", err)
	}
}
