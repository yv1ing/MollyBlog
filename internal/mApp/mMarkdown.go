package mApp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"MollyBlog/internal/model"
	"MollyBlog/utils"

	"gopkg.in/yaml.v3"
)

// loadMarkdownFiles load markdown source files
func (ma *MApp) loadMarkdownFiles() error {
	var markdownList []model.MFileInfo

	markdownPath := SRC
	err := filepath.Walk(markdownPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// filter markdown files
		if filepath.Ext(path) == ".md" {
			markdownList = append(markdownList, model.MFileInfo{
				Name: info.Name(),
				Path: path,
			})
		}

		return nil
	})

	if err != nil {
		return err
	}

	ma.SrcFiles = markdownList
	return nil
}

// parseMarkdowns parse markdown files to html
func (ma *MApp) parseMarkdowns() error {
	htmlPath := DST

	for _, file := range ma.SrcFiles {
		// read markdown file
		_mdFile, err := os.Open(file.Path)
		if err != nil {
			return err
		}
		defer _mdFile.Close()

		_mdByte, err := io.ReadAll(_mdFile)
		if err != nil {
			return err
		}

		// extract FrontMatter and clean markdown bytes
		_frontMatter, _cleanBytes := utils.ExtractFrontMatter(_mdByte)

		// parse frontmatter to post metadata
		var post model.MPost
		err = yaml.Unmarshal(_frontMatter, &post)
		if err != nil {
			return err
		}

		// convert into html format and save to html file
		_htmlPath := fmt.Sprintf("%s/%s.html", htmlPath, file.Name)
		_htmlFile, err := os.Create(_htmlPath)
		if err != nil {
			return err
		}
		defer _htmlFile.Close()

		// set options to lute
		ma.lute.SetToC(true)

		_htmlByte := ma.lute.Markdown(file.Name, _cleanBytes)
		_, err = _htmlFile.Write(_htmlByte)
		if err != nil {
			return err
		}

		// save html path to post's metadata
		post.HtmlHash = utils.Sha256Hash(_htmlByte)
		post.HtmlPath = _htmlPath

		// save tags and categories to map
		for _, tag := range post.Tags {
			tagHash := utils.Sha256Hash([]byte(tag))
			post.TagHashes = append(post.TagHashes, model.MTag{
				Name: tag,
				Hash: tagHash,
			})

			ma.TaggedPosts = append(ma.TaggedPosts, &post)
		}

		for _, category := range post.Categories {
			categoryHash := utils.Sha256Hash([]byte(category))
			post.CategoryHashes = append(post.CategoryHashes, model.MCategory{
				Name: category,
				Hash: categoryHash,
			})

			ma.CategorizedPosts = append(ma.CategorizedPosts, &post)
		}

		// free the raw tag and category slice
		post.Tags = nil
		post.Categories = nil

		ma.Posts = append(ma.Posts, &post)
	}

	// sort Posts by date
	model.SortPostsByDate(ma.Posts)
	return nil
}
