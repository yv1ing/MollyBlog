package mApp

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"MollyBlog/internal/model"
	"MollyBlog/utils"

	"github.com/gin-gonic/gin"
	"github.com/huichen/wukong/types"
)

func (ma *MApp) IndexHandler(ctx *gin.Context) {
	// generate recent posts
	var recentPosts []model.MPost
	for i := 0; i < utils.Min(len(ma.Posts), ma.Config.MSite.Post.RecentPost.Number); i++ {
		tmpPost := *ma.Posts[i]
		tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]

		recentPosts = append(recentPosts, tmpPost)
	}

	// generate motto
	var mottoHtml string
	mottoHtml = "<script>let motto = ['"
	for _, motto := range ma.Config.MSite.Info.Motto {
		mottoHtml += fmt.Sprintf("%s<br>", motto)
	}
	mottoHtml += "'];</script>"

	// return some basic information
	resData := gin.H{
		"site_info": gin.H{
			"logo":       ma.Config.MSite.Info.Logo,
			"title":      ma.Config.MSite.Info.Title,
			"author":     ma.Config.MSite.Info.Author,
			"language":   ma.Config.MSite.Info.Language,
			"copyright":  template.HTML(ma.Config.MSite.Info.Copyright),
			"motto_html": template.HTML(mottoHtml),
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

	// traverse to find the corresponding post, read its HTML file, and inject it into the template
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

func (ma *MApp) TagHandler(ctx *gin.Context) {
	tagHash := ctx.Param("hash")
	tagName := ma.Tags[tagHash]

	// paging logic processing
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size := ma.Config.MSite.Post.Tag.Number

	var prePage, curPage, nxtPage, allPage int
	allPage = (len(ma.TaggedPosts[tagHash]) + size - 1) / size

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

	// generate tagged posts
	start := (curPage - 1) * size
	offset := curPage * size

	var taggedPosts []model.MPost
	var tagList [][]interface{}
	if start >= 0 {
		for i := start; i < utils.Min(len(ma.TaggedPosts[tagHash]), offset); i++ {
			tmpPost := *ma.TaggedPosts[tagHash][i]
			tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]
			taggedPosts = append(taggedPosts, tmpPost)
		}

		for tag, num := range ma.TagsCount {
			tagList = append(tagList, []interface{}{tag, num, ma.TagsHash[tag]})
		}
	}

	tagListJson, _ := json.Marshal(tagList)
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
		"tagged_post": gin.H{
			"title":    fmt.Sprintf("%s - %s", ma.Config.MSite.Post.Tag.Title, tagName),
			"posts":    taggedPosts,
			"tag_hash": tagHash,
			"tag_list": string(tagListJson),
		},
	}

	ctx.HTML(http.StatusOK, "tag.html", resData)
}

func (ma *MApp) CategoryHandler(ctx *gin.Context) {
	categoryHash := ctx.Param("hash")
	categoryName := ma.Categories[categoryHash]

	// paging logic processing
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size := ma.Config.MSite.Post.Tag.Number

	var prePage, curPage, nxtPage, allPage int
	allPage = (len(ma.CategorizedPosts[categoryHash]) + size - 1) / size

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

	// generate categorized posts
	start := (curPage - 1) * size
	offset := curPage * size

	var categorizedPosts []model.MPost
	if start >= 0 {
		for i := start; i < utils.Min(len(ma.CategorizedPosts[categoryHash]), offset); i++ {
			tmpPost := *ma.CategorizedPosts[categoryHash][i]
			tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]
			categorizedPosts = append(categorizedPosts, tmpPost)
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
		"categorized_post": gin.H{
			"title":         fmt.Sprintf("%s - %s", ma.Config.MSite.Post.Category.Title, categoryName),
			"posts":         categorizedPosts,
			"category_hash": categoryHash,
		},
	}

	ctx.HTML(http.StatusOK, "category.html", resData)
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

	// generate history posts
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

	// generate motto
	var mottoHtml string
	mottoHtml = "<script>let motto = ['"
	for _, motto := range ma.Config.MSite.Info.Motto {
		mottoHtml += fmt.Sprintf("%s<br>", motto)
	}
	mottoHtml += "'];</script>"

	resData := gin.H{
		"site_info": gin.H{
			"logo":       ma.Config.MSite.Info.Logo,
			"title":      ma.Config.MSite.Info.Title,
			"author":     ma.Config.MSite.Info.Author,
			"language":   ma.Config.MSite.Info.Language,
			"copyright":  template.HTML(ma.Config.MSite.Info.Copyright),
			"motto_html": template.HTML(mottoHtml),
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

func (ma *MApp) SearchHandler(ctx *gin.Context) {
	keyword := ctx.DefaultQuery("keyword", "")
	if keyword == "" || len(ma.Posts) <= 1 {
		ctx.Redirect(http.StatusFound, "/archive")
		return
	}

	searchResult := ma.searcher.Search(types.SearchRequest{Text: keyword})
	searchPosts := searchResult.Docs

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size := ma.Config.MSite.Post.Search.Number

	var prePage, curPage, nxtPage, allPage int
	allPage = (len(searchPosts) + size - 1) / size

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

	// generate search posts
	start := (curPage - 1) * size
	offset := curPage * size

	var resultPosts []model.MPost
	if len(searchPosts) > 0 {
		for i := start; i < utils.Min(len(searchPosts), offset); i++ {
			resultPosts = append(resultPosts, *ma.IndexedPosts[searchPosts[i].DocId])
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
		"search_post": gin.H{
			"title":   fmt.Sprintf("%s - %s", ma.Config.MSite.Post.Search.Title, keyword),
			"posts":   resultPosts,
			"keyword": keyword,
		},
	}

	ctx.HTML(http.StatusOK, "search.html", resData)
}

func (ma *MApp) UpdateBlogHandler(ctx *gin.Context) {
	var err error
	err = ma.resetStorage()
	if err != nil {
		log.Printf("reset storage error: %v\n", err)
		_ = ctx.Error(err)
		return
	}

	// load posts
	err = ma.loadMarkdownFiles()
	if err != nil {
		log.Printf("load markdown files failed, err: %v\n", err)
		_ = ctx.Error(err)
		return
	}

	err = ma.parseMarkdowns()
	if err != nil {
		log.Printf("parse markdown files failed, err: %v\n", err)
		_ = ctx.Error(err)
		return
	}

	// load about me
	aboutSrc := fmt.Sprintf("%s/%s.md", ma.Config.MSite.About.SRC, ma.Config.MSite.About.Filename)
	aboutDst := fmt.Sprintf("%s/%s.html", ma.Config.MSite.About.DST, ma.Config.MSite.About.Filename)
	err = ma.parseSingleMarkdown(aboutSrc, aboutDst)
	if err != nil {
		log.Printf("parse about me failed, err: %v\n", err)
		_ = ctx.Error(err)
		return
	}

	// parse post index
	ma.loadPostIndex()

	// parse rss
	ma.RSS = ma.generateRSS()

	log.Println("update blog success")
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (ma *MApp) AboutHandler(ctx *gin.Context) {
	dstFile := fmt.Sprintf("%s/%s.html", ma.Config.MSite.About.DST, ma.Config.MSite.About.Filename)
	aboutFile, _ := os.OpenFile(dstFile, os.O_RDONLY, 0666)
	aboutBytes, _ := io.ReadAll(aboutFile)

	// return some basic information
	resData := gin.H{
		"site_info": gin.H{
			"logo":      ma.Config.MSite.Info.Logo,
			"title":     ma.Config.MSite.Info.Title,
			"author":    ma.Config.MSite.Info.Author,
			"language":  ma.Config.MSite.Info.Language,
			"copyright": template.HTML(ma.Config.MSite.Info.Copyright),
		},
		"menu": ma.Config.MSite.Menu,
		"about": gin.H{
			"title":   ma.Config.MSite.About.Title,
			"content": template.HTML(aboutBytes),
			"success": true,
		},
	}

	ctx.HTML(http.StatusOK, "about.html", resData)
}

func (ma *MApp) RSSHandler(ctx *gin.Context) {
	if ma.RSS != "" {
		ctx.String(http.StatusOK, ma.RSS)
	} else {
		ctx.String(http.StatusNotFound, "RSS Not Found")
	}
}

func (ma *MApp) FriendHandler(ctx *gin.Context) {
	resData := gin.H{
		"site_info": gin.H{
			"logo":      ma.Config.MSite.Info.Logo,
			"title":     ma.Config.MSite.Info.Title,
			"author":    ma.Config.MSite.Info.Author,
			"language":  ma.Config.MSite.Info.Language,
			"copyright": template.HTML(ma.Config.MSite.Info.Copyright),
		},
		"menu": ma.Config.MSite.Menu,
		"friend": gin.H{
			"title": ma.Config.MSite.Friend.Title,
			"list":  ma.Config.MSite.Friend.List,
		},
	}
	ctx.HTML(http.StatusOK, "friend.html", resData)
}
