package mApp

import (
	"io"
	"os"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

func (ma *MApp) loadPostIndex() {
	ma.searcher = &engine.Engine{}

	ma.searcher.Init(types.EngineInitOptions{
		UsePersistentStorage:    true,
		PersistentStorageFolder: "tmp",
		StopTokenFile:           "data/stop_tokens.txt",
		SegmenterDictionaries:   "data/dictionary.txt",
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.LocationsIndex,
		},
	})

	for _, post := range ma.Posts {
		postFile, _ := os.OpenFile(post.HtmlPath, os.O_RDONLY, 0666)
		postData, _ := io.ReadAll(postFile)

		ma.searcher.IndexDocument(post.Index, types.DocumentIndexData{Content: string(postData)}, false)
	}

	ma.searcher.FlushIndex()
}
