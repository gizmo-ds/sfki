package main

import (
	"io/ioutil"
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
	r := chi.NewRouter()
	r.Get("/", control.Home)
	http.ListenAndServe(config.Addr, r)
}
