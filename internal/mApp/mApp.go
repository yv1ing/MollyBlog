package mApp

import (
	"fmt"
	"html/template"
	"log"

	"MollyBlog/config"
	"MollyBlog/internal/model"

	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
)

type MApp struct {
	Host   string
	Port   int
	Config *config.MConfig

	lute   *lute.Lute
	engine *gin.Engine

	Posts            []*model.MPost
	TaggedPosts      []*model.MPost
	CategorizedPosts []*model.MPost

	SrcFiles []model.MFileInfo
}

const (
	SRC = "_post/src" // source markdown files
	DST = "_post/dst" // destination html files
)

func (ma *MApp) Run() {
	ma.loadRoutes()
	ma.loadTemplates()

	addr := fmt.Sprintf("%s:%d", ma.Host, ma.Port)
	err := ma.engine.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func NewMApp(cfg *config.MConfig) *MApp {
	engine := gin.Default()
	engine.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	return &MApp{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Config: cfg,

		lute:   lute.New(),
		engine: engine,
	}
}
