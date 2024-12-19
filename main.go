package main

import (
	"MollyBlog/config"
	"MollyBlog/internal/mApp"
	"MollyBlog/utils"
)

func init() {
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
	mapp.Run(true)
}
