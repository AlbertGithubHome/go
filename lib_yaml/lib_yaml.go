package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
    APIKey string `yaml:"api_key"`
    Port   int    `yaml:"port"`
}

func main() {
    config := Configuration{
        APIKey: "secret-key",
        Port:   8080,
    }

    yamlData, _ := yaml.Marshal(config)
    fmt.Println("YAML Data:")
    fmt.Println(string(yamlData))
}

/*
$ go run .
YAML Data:
api_key: secret-key
port: 8080
*/