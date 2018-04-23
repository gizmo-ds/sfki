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
		Context string   `json:"context"`
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
		"context": &graphql.Field{
			Type: graphql.String,
		},
	},
})
