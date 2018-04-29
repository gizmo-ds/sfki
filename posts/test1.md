<!--
title: 测试文章1
alias: test1
created: 2018-04-19
updated: 2018-04-19
tags: [test1,test2]
-->

## Test1  
> Hello  
<!-- More -->

|id|name|
|-|-|
|1|A1|
|2|A2|
|3|A3|

``` go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(model.ExecuteQuery(`{posts(tag:"test1"){alias,title}}`))
	})
	r.Get("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("query")
		json.NewEncoder(w).Encode(model.ExecuteQuery(query))
	})
	r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
		model.PostLoading()
	})
	http.ListenAndServe(config.Addr, r)
}
```
