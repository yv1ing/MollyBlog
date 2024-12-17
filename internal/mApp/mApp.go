package mApp

import (
	"fmt"
	"html/template"
	"log"

	"MollyBlog/config"
	"MollyBlog/internal/model"

	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"github.com/huichen/wukong/engine"
)

type MApp struct {
	Host   string
	Port   int
	Config *config.MConfig

	lute     *lute.Lute
	engine   *gin.Engine
	searcher *engine.Engine

	Posts        []*model.MPost
	IndexedPosts map[uint64]*model.MPost

	Tags             map[string]string
	TagsCount        map[string]int
	Categories       map[string]string
	CategoriesCount  map[string]int
	TaggedPosts      map[string][]*model.MPost
	CategorizedPosts map[string][]*model.MPost

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

		IndexedPosts: make(map[uint64]*model.MPost),

		Tags:      make(map[string]string),
		TagsCount: make(map[string]int),

		Categories:      make(map[string]string),
		CategoriesCount: make(map[string]int),

		TaggedPosts:      make(map[string][]*model.MPost),
		CategorizedPosts: make(map[string][]*model.MPost),

		lute:   lute.New(),
		engine: engine,
	}
}

// resetStorage before each update, delete the cache
func (ma *MApp) resetStorage() {
	ma.Posts = nil
	ma.Tags = nil
	ma.Categories = nil
	ma.CategoriesCount = nil
	ma.TaggedPosts = nil
	ma.CategorizedPosts = nil
	ma.SrcFiles = nil

	ma.Tags = make(map[string]string)
	ma.TagsCount = make(map[string]int)
	ma.Categories = make(map[string]string)
	ma.CategoriesCount = make(map[string]int)
	ma.TaggedPosts = make(map[string][]*model.MPost)
	ma.CategorizedPosts = make(map[string][]*model.MPost)
}
