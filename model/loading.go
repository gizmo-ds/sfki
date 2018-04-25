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
	ROOT   string
	Posts  sync.Map
	TabMap sync.Map
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

	PostLoading()
}

// 读取Post
func PostLoading() {
	Posts = sync.Map{}
	TabMap = sync.Map{}

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

		_info, ok := Posts.Load(info.Alias)
		if ok {
			return errors.New(
				fmt.Sprintf("Alias Exist: \n%v\n%v", _info.(Post).Path, info.Path))
		}
		Posts.Store(info.Alias, info)

		for _, v := range info.Tags {
			_post_old, ok := TabMap.Load(v)
			if !ok {
				_post_old = []Post{info}
			} else {
				// TODO: 有可能会boom
				_post_old = append(_post_old.([]Post), info)
			}
			TabMap.Store(v, _post_old)
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
