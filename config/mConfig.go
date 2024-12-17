package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MConfig struct {
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	Storage        mStorage `yaml:"storage"`
	Template       string   `yaml:"template"`
	UpdateEndpoint string   `yaml:"update_endpoint"`
	UpdateSecret   string   `yaml:"update_secret"`
	CachePath      string   `yaml:"cache_path"`

	MSite mSite `yaml:"site"`
}

var MConfigInstance *MConfig

func init() {
	configFile, err := os.OpenFile("config.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configByte, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configByte, &MConfigInstance)
	if err != nil {
		log.Fatal(err)
	}
}
