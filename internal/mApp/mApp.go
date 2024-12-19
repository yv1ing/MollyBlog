package mApp

import (
	"fmt"
	"html/template"
	"log"
	"os"

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

	RSS string

	Posts        []*model.MPost
	IndexedPosts map[uint64]*model.MPost

	Tags             map[string]string
	TagsHash         map[string]string
	TagsCount        map[string]int
	Categories       map[string]string
	CategoriesCount  map[string]int
	TaggedPosts      map[string][]*model.MPost
	CategorizedPosts map[string][]*model.MPost

	SrcFiles []model.MFileInfo
}

func init() {
	log.SetPrefix("[MollyBlog] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	gin.SetMode(gin.ReleaseMode)
}

func (ma *MApp) Run() {
	ma.loadRoutes()
	ma.loadTemplates()

	addr := fmt.Sprintf("%s:%d", ma.Host, ma.Port)
	log.Printf("app listening on %s\n", addr)
	err := ma.engine.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func NewMApp(cfg *config.MConfig) *MApp {
	_engine := gin.Default()
	_engine.SetFuncMap(template.FuncMap{
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
		TagsHash:  make(map[string]string),
		TagsCount: make(map[string]int),

		Categories:      make(map[string]string),
		CategoriesCount: make(map[string]int),

		TaggedPosts:      make(map[string][]*model.MPost),
		CategorizedPosts: make(map[string][]*model.MPost),

		lute:   lute.New(),
		engine: _engine,
	}
}

// resetStorage before each update, delete the cache
func (ma *MApp) resetStorage() error {
	ma.Posts = nil
	ma.Tags = nil
	ma.TagsHash = nil
	ma.TagsCount = nil
	ma.Categories = nil
	ma.CategoriesCount = nil
	ma.TaggedPosts = nil
	ma.CategorizedPosts = nil
	ma.SrcFiles = nil

	ma.Tags = make(map[string]string)
	ma.TagsHash = make(map[string]string)
	ma.TagsCount = make(map[string]int)
	ma.Categories = make(map[string]string)
	ma.CategoriesCount = make(map[string]int)
	ma.TaggedPosts = make(map[string][]*model.MPost)
	ma.CategorizedPosts = make(map[string][]*model.MPost)

	if ma.searcher != nil {
		ma.searcher.Close()
	}

	var err error
	if ma.Config.Storage.AutoPull {
		err = os.RemoveAll(ma.Config.Storage.SRC)
		if err != nil {
			return err
		}

		_, err = os.Stat(ma.Config.Storage.SRC)
		if os.IsNotExist(err) {
			err = os.MkdirAll(ma.Config.Storage.SRC, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	err = os.RemoveAll(ma.Config.Storage.DST)
	if err != nil {
		return err
	}

	_, err = os.Stat(ma.Config.Storage.DST)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ma.Config.Storage.DST, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.RemoveAll(ma.Config.CachePath)
	if err != nil {
		return err
	}

	return nil
}
