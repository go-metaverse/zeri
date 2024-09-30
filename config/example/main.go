package main

import (
	"fmt"

	"github.com/go-metaverse/zeri/config"
)

type App struct {
	Name      string   `json:"name" yaml:"name"`
	Version   string   `json:"version" yaml:"version"`
	Databases Database `json:"database" yaml:"database"`
}

type Database struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

func main() {
	appConfig, err := config.NewConfig(config.GetConfigPath("local", ".json"), &App{})
	// appConfig, err := config.NewConfig("./config/example/env.sample.yml", &App{})
	if err != nil {
		panic(err)
	}

	fmt.Println(appConfig)
}
