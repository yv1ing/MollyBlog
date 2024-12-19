package main

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"

	"MollyBlog/config"
	"MollyBlog/internal/mApp"
	"MollyBlog/utils"
)

func init() {
	// parse arguments
	configFilePath := flag.String("c", "config.yaml", "config file")

	flag.Parse()

	// load config
	configFile, err := os.OpenFile(*configFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configByte, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configByte, &config.MConfigInstance)
	if err != nil {
		log.Fatal(err)
	}

	// create need files
	postSrc := config.MConfigInstance.Storage.SRC
	postDst := config.MConfigInstance.Storage.DST
	utils.Mkdir(postSrc)
	utils.Mkdir(postDst)

	aboutSrc := config.MConfigInstance.MSite.About.SRC
	aboutDst := config.MConfigInstance.MSite.About.DST
	utils.Mkdir(aboutSrc)
	utils.Mkdir(aboutDst)
}

func main() {
	mapp := mApp.NewMApp(config.MConfigInstance)
	mapp.Run(false)
}
