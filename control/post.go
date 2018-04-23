package control

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/loadfield/sfki/model"
)

type H map[string]interface{}

func Tags(w http.ResponseWriter, r *http.Request) {
	type tag struct {
		Name  string
		Count int
	}
	_tags := []tag{}
	for k, v := range model.TagMap {
		_tags = append(_tags, tag{Name: k, Count: len(v)})
	}
	Json(w, H{"success": true, "tags": _tags})
}

func TagPosts(w http.ResponseWriter, r *http.Request) {
	tag_name := chi.URLParam(r, "tag_name")
	Json(w, H{"success": true, "posts": model.TagMap[tag_name]})
}

func Posts(w http.ResponseWriter, r *http.Request) {
	_posts := []model.Post{}
	model.Posts.Range(func(k, v interface{}) bool {
		// TODO: 这个倒序插入可能不靠谱
		_posts = append([]model.Post{v.(model.Post)}, _posts...)
		return true
	})
	Json(w, H{"success": true, "posts": _posts})
}

func Json(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	bytes, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"success":false,"error":%v}`, err.Error())))
		return
	}
	w.Write(bytes)
}
