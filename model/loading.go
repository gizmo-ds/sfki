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

type (
	post struct {
		Path    string `yaml:"-"`
		Title   string
		Alias   string
		Created string
		Updated string
		Tags    []string
	}
)

var (
	ROOT   string
	Posts  sync.Map
	TagMap map[string][]post
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
	TagMap = make(map[string][]post)
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
		var info post
		if err = yaml.Unmarshal(bytes[4:index], &info); err != nil {
			return errors.New(
				fmt.Sprintf("YAML Unmarshal Error: %v (%v)", err.Error(), path))
		}
		info.Path = path

		_info, ok := Posts.Load(info.Alias)
		if ok {
			return errors.New(
				fmt.Sprintf("Alias Exist: \n%v\n%v", _info.(post).Path, info.Path))
		}
		Posts.Store(info.Alias, info)
		// Posts = append(Posts, info)

		for _, v := range info.Tags {
			TagMap[v] = append(TagMap[v], info)
		}
		// log.Println(string(bytes[index+4:]))
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
	// log.Println(TagMap)
	// log.Println(Posts)
}
