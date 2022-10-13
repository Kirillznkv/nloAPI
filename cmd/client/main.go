package main

import (
	"flag"
	"github.com/Kirillznkv/nloAPI/internal/app/client"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Kirillznkv/nloAPI/internal/pkg/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/store.toml", "path to toml config file")
}

func main() {
	flag.Parse()

	config := store.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	c := client.New(config)
	if err := c.Start(); err != nil {
		log.Fatal(err)
	}
}
