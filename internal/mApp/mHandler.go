package mApp

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"MollyBlog/internal/model"
	"MollyBlog/utils"

	"github.com/gin-gonic/gin"
)

func (ma *MApp) IndexHandler(ctx *gin.Context) {
	// generate recent posts
	var recentPosts []model.MPost
	for i := 0; i < utils.Min(len(ma.Posts), ma.Config.MSite.Post.RecentPost.Number); i++ {
		tmpPost := *ma.Posts[i]
		tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]

		recentPosts = append(recentPosts, tmpPost)
	}

	resData := gin.H{
		"site_info": gin.H{
			"logo":      ma.Config.MSite.Info.Logo,
			"title":     ma.Config.MSite.Info.Title,
			"author":    ma.Config.MSite.Info.Author,
			"language":  ma.Config.MSite.Info.Language,
			"copyright": template.HTML(ma.Config.MSite.Info.Copyright),
		},
		"menu": ma.Config.MSite.Menu,
		"recent_post": gin.H{
			"title": ma.Config.MSite.Post.RecentPost.Title,
			"posts": recentPosts,
		},
	}

	ctx.HTML(http.StatusOK, "index.html", resData)
}

func (ma *MApp) PostHandler(ctx *gin.Context) {
	postHash := ctx.Param("hash")

	var success bool
	var html string
	var realPost model.MPost
	for _, post := range ma.Posts {
		if post.HtmlHash == postHash {
			file, err := os.OpenFile(post.HtmlPath, os.O_RDONLY, 0644)
			if err != nil {
				success = false
				break
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				success = false
				break
			}

			html = string(data)
			realPost = *post
			success = true
			break
		}
	}

	var resPost gin.H
	if success {
		resPost = gin.H{
			"title":      realPost.Title,
			"cover":      realPost.Cover,
			"date":       realPost.Date,
			"tags":       realPost.TagHashes,
			"categories": realPost.CategoryHashes,
			"content":    template.HTML(html),
		}
	}

	resData := gin.H{
		"site_info": gin.H{
			"logo":      ma.Config.MSite.Info.Logo,
			"title":     ma.Config.MSite.Info.Title,
			"author":    ma.Config.MSite.Info.Author,
			"language":  ma.Config.MSite.Info.Language,
			"copyright": template.HTML(ma.Config.MSite.Info.Copyright),
		},
		"menu": ma.Config.MSite.Menu,
		"post": resPost,
		"status": gin.H{
			"toc_title": ma.Config.MSite.Post.TocTitle,
			"success":   success,
		},
	}

	ctx.HTML(http.StatusOK, "post.html", resData)
}

func (ma *MApp) ArchiveHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size := ma.Config.MSite.Post.Archive.Number

	var prePage, curPage, nxtPage, allPage int
	allPage = (len(ma.Posts) + size - 1) / size

	if allPage > 0 {
		if page <= 0 {
			curPage = 1
		} else if page > allPage {
			curPage = allPage
		} else {
			curPage = page
		}
	} else {
		curPage = 0
	}

	prePage = curPage - 1
	nxtPage = curPage + 1

	if prePage <= 0 {
		prePage = curPage
	}

	if nxtPage > allPage {
		nxtPage = allPage
	}

	// generate recent posts
	start := (curPage - 1) * size
	offset := curPage * size

	var historyPosts []model.MPost
	if start >= 0 {
		for i := start; i < utils.Min(len(ma.Posts), offset); i++ {
			tmpPost := *ma.Posts[i]
			tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]
			historyPosts = append(historyPosts, tmpPost)
		}
	}

	resData := gin.H{
		"site_info": gin.H{
			"logo":      ma.Config.MSite.Info.Logo,
			"title":     ma.Config.MSite.Info.Title,
			"author":    ma.Config.MSite.Info.Author,
			"language":  ma.Config.MSite.Info.Language,
			"copyright": template.HTML(ma.Config.MSite.Info.Copyright),
		},
		"menu": ma.Config.MSite.Menu,
		"page_info": gin.H{
			"pre_page": prePage,
			"cur_page": curPage,
			"nxt_page": nxtPage,
			"all_page": allPage,
		},
		"history_post": gin.H{
			"title": ma.Config.MSite.Post.Archive.Title,
			"posts": historyPosts,
		},
	}

	ctx.HTML(http.StatusOK, "archive.html", resData)
}

func (ma *MApp) UpdateBlogHandler(ctx *gin.Context) {
	var err error

	err = ma.loadMarkdownFiles()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = ma.parseMarkdowns()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
