package main

import (
	"log"
	"os"

	"MollyBlog/config"
	"MollyBlog/internal/mApp"
)

func init() {
	src := config.MConfigInstance.Storage.SRC
	dst := config.MConfigInstance.Storage.DST

	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		err = os.MkdirAll(src, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stat(dst)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	mapp := mApp.NewMApp(config.MConfigInstance)
	mapp.Run()
}
