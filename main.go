package main

import (
	"fmt"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/api"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
)

func main() {
	err := config.ReadConfig("configuration.yml")
	if err != nil {
		fmt.Printf("Error while reading config file: %v\n", err)
		return
	}
	api.StartServer()
}
