package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	ROOT  string
	Post_ struct {
		sync.Mutex
		posts []Post
	}
	Link_ struct {
		sync.Mutex
		links []Link
	}
	TagMap sync.Map
)

func init() {
	ROOT = os.Getenv("SFKI_ROOT")
	if ROOT == "" {
		var err error
		ROOT, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	LinkLoading()
	PostLoading()
}

func LinkLoading() {
	Link_.Lock()
	defer func() {
		Link_.Unlock()
	}()
	Link_.links = []Link{}

	bytes, err := ioutil.ReadFile(filepath.Join(ROOT, "links.yaml"))
	if err == nil {
		if err = yaml.Unmarshal(bytes, &Link_.links); err != nil {
			panic(err)
		}
	}
}

func PostLoading() {
	Post_.Lock()
	defer func() {
		Post_.Unlock()
	}()
	Post_.posts = []Post{}
	TagMap = sync.Map{}

	_load := func(path string) error {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		index := strings.Index(string(bytes), "-->\n")
		if index == -1 {
			return errors.New(
				fmt.Sprintf("Get Info Error: '-->\\n' Not Exist (%v)", path))
		}
		var info Post
		if err = yaml.Unmarshal(bytes[4:index], &info); err != nil {
			return errors.New(
				fmt.Sprintf("YAML Unmarshal Error: %v (%v)", err.Error(), path))
		}
		info.Path = path
		info.Content = string(bytes)
		Post_.posts = append(Post_.posts, info)

		for _, v := range info.Tags {
			_post_old, ok := TagMap.Load(v)
			if !ok {
				_post_old = []Post{info}
			} else {
				// TODO: 有可能会boom
				_post_old = append(_post_old.([]Post), info)
			}
			TagMap.Store(v, _post_old)
		}
		return nil
	}

	filepath.Walk(filepath.Join(ROOT, "posts/"),
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			if err := _load(path); err != nil {
				log.New(os.Stdout, "[Warning] model.loading.PostLoading() ",
					log.LstdFlags).Println(err.Error())
			}
			return nil
		})
}
