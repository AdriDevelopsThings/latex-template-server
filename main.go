package main

import (
	"fmt"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/api"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/files"
)

func main() {
	err := config.ReadConfig("configuration.yml")
	if err != nil {
		fmt.Printf("Error while reading config file: %v\n", err)
		return
	}
	go files.StartAutoDeleteFiles()
	api.StartServer()
}
