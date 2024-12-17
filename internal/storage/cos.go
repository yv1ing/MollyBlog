package storage

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"MollyBlog/config"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func cosDownload(wg *sync.WaitGroup, c *cos.Client, keysCh <-chan []string) {
	defer wg.Done()
	for keys := range keysCh {
		key := keys[0]
		filename := keys[1]
		_, err := c.Object.GetToFile(context.Background(), key, filename, nil)
		if err != nil {
			log.Println(err)
		}

		log.Println("download success: ", filename)
	}
}

func CosLoadMarkdowns(config config.MConfig, localDir string) error {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", config.Storage.COS.Bucket, config.Storage.COS.Region))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Storage.COS.SecretId,
			SecretKey: config.Storage.COS.SecretKey,
		},
	})

	var wg sync.WaitGroup

	keysCh := make(chan []string, 3)
	threadPoolNum := 3
	for i := 0; i < threadPoolNum; i++ {
		wg.Add(1)
		go cosDownload(&wg, c, keysCh)
	}

	marker := ""
	prefix := config.Storage.COS.SavePath
	encodingType := "url"
	
	isTruncated := true
	for isTruncated {
		opt := &cos.BucketGetOptions{
			Prefix:       prefix,
			Marker:       marker,
			EncodingType: encodingType,
		}

		v, _, err := c.Bucket.Get(context.Background(), opt)
		if err != nil {
			return err
		}

		for _, c := range v.Contents {
			key, _ := cos.DecodeURIComponent(c.Key)
			localFile := path.Join(localDir, key)
			if _, err := os.Stat(path.Dir(localFile)); err != nil && os.IsNotExist(err) {
				_ = os.MkdirAll(path.Dir(localFile), os.ModePerm)
			}

			if strings.HasSuffix(localFile, "/") {
				continue
			}

			keysCh <- []string{key, localFile}
		}
		marker, _ = cos.DecodeURIComponent(v.NextMarker)
		isTruncated = v.IsTruncated
	}

	close(keysCh)
	wg.Wait()

	return nil
}
