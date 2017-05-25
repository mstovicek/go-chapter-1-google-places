package google_api

import (
	"fmt"
	"github.com/joeshaw/envdecode"
)

type Config struct {
	ApiKey string `env:"API_KEY"`
}

func Ping() {
	var config Config
	err := envdecode.Decode(&config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.ApiKey)
}
