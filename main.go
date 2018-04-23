package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"

	"github.com/loadfield/sfki/control"
	"github.com/loadfield/sfki/model"

	"github.com/go-chi/chi"

	_ "github.com/loadfield/sfki/model"
)

var (
	config struct {
		Addr string
	}
)

func init() {
	bytes, err := ioutil.ReadFile(model.ROOT + "/config.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}
}

func main() {

	log.Println(model.ExecuteQuery(`{posts(tag:"test1"){alias,title}}`))

	r := chi.NewRouter()
	r.Get("/tags", control.Tags)
	r.Get("/tag/{tag_name}", control.TagPosts)
	r.Get("/posts", control.Posts)
	http.ListenAndServe(config.Addr, r)
}
