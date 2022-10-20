package config

import (
	"fmt"
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
	Salt              string `yaml:"salt"`
	EncryptionKeySize int    `yaml:"encryption_key_size"`
	FileServePath     string `yaml:"file_serve_path"`
	TemplatePath      string `yaml:"template_path"`
	TmpDirectory      string `yaml:"tmp_directory"`
	AppUrl            string `yaml:"app_url"`
	DeleteFileAfter   int    `yaml:"delete_file_after"`
}

func setupDefaults() {
	CurrentConfig.Server.Listen.Port = "3000"
	CurrentConfig.EncryptionKeySize = 48
	CurrentConfig.FileServePath = "files"
	CurrentConfig.TemplatePath = "templates"
	CurrentConfig.TmpDirectory = "tmp"
	CurrentConfig.AppUrl = "http://localhost:3000"
	CurrentConfig.DeleteFileAfter = 60 * 10 // 10 minutes
}

func validateConfig() error {
	if CurrentConfig.Salt == "" {
		return fmt.Errorf("no_salt_set_in_config", "No salt is set in config file: Generate a hex salt with openssl rand -hex 32")
	}
	os.Mkdir(CurrentConfig.FileServePath, os.ModePerm)
	os.Mkdir(CurrentConfig.TemplatePath, os.ModePerm)
	os.Mkdir(CurrentConfig.TmpDirectory, os.ModePerm)

	return nil
}

func ReadConfig(filepath string) error {
	setupDefaults()
	config, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(config, &CurrentConfig)
	if err != nil {
		return err
	}
	return validateConfig()
}
