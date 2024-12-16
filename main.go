package main

import (
	"log"
	"os"

	"MollyBlog/config"
	"MollyBlog/internal/mApp"
)

func init() {
	_, err := os.Stat(mApp.SRC)
	if os.IsNotExist(err) {
		err = os.MkdirAll(mApp.SRC, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stat(mApp.DST)
	if os.IsNotExist(err) {
		err = os.MkdirAll(mApp.DST, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	mapp := mApp.NewMApp(config.MConfigInstance)
	mapp.Run()
}
