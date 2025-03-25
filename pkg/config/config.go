package config

import (
	"errors"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

var CurrentConfig Config

type Config struct {
	Server struct {
		Listen struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"listen"`
	} `yaml:"server"`
	TemplatePath string `yaml:"template_path"`
	TmpDirectory string `yaml:"tmp_directory"`
}

func setupDefaults() {
	CurrentConfig.Server.Listen.Port = "3000"
	CurrentConfig.TemplatePath = "templates"
	CurrentConfig.TmpDirectory = "tmp"
}

func validateConfig() error {
	os.Mkdir(CurrentConfig.TemplatePath, os.ModePerm)
	os.Mkdir(CurrentConfig.TmpDirectory, os.ModePerm)

	return nil
}

func ReadConfig(filepath string) error {
	setupDefaults()
	_, err := os.Stat(filepath)
	if err == nil {
		config, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(config, &CurrentConfig)
		if err != nil {
			return err
		}
	} else {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}
	return validateConfig()
}
