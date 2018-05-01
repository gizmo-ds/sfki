package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"

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
	r.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.PostFormValue("query")
		json.NewEncoder(w).Encode(model.ExecuteQuery(query))
	})
	r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
		model.PostLoading()
		model.LinkLoading()
	})
	http.ListenAndServe(config.Addr, r)
}
