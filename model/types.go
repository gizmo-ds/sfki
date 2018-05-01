package model

import (
	"github.com/graphql-go/graphql"
)

type (
	Post struct {
		Path    string   `json:"-"`
		Title   string   `json:"title"`
		Alias   string   `json:"alias"`
		Created string   `json:"created"`
		Updated string   `json:"updated"`
		Tags    []string `json:"tags"`
		Content string   `json:"content"`
	}
	Link struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
	}
)

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"alias": &graphql.Field{
			Type: graphql.String,
		},
		"created": &graphql.Field{
			Type: graphql.String,
		},
		"updated": &graphql.Field{
			Type: graphql.String,
		},
		"tags": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var tagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "tag",
	Fields: graphql.Fields{
		"tag": &graphql.Field{
			Type: graphql.String,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(postType),
		},
	},
})

var postsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"max": &graphql.Field{
			Type: graphql.Int,
		},
		"list": &graphql.Field{
			Type: graphql.NewList(postType),
		},
	},
})

var linkType = graphql.NewObject(graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"link": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})
