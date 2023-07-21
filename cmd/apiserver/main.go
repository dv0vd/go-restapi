package main

import (
	"log"
	"os"

	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/apiserver"
)

func init() {
	os.Setenv("APISERVER_CONFIG_PATH", "configs/apiserver.toml")
}

func main() {
	apiServerConfig, err := apiserver.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(apiServerConfig); err != nil {
		log.Fatal(err)
	}
}
