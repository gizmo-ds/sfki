package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: postType,
			Args: graphql.FieldConfigArgument{
				"alias": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_v := Post{}

				aliasQ, ok := params.Args["alias"].(string)
				if ok {
					Posts.Range(func(_, v interface{}) bool {
						_v = v.(Post)
						if _v.Alias == aliasQ {
							bytes, err := ioutil.ReadFile(_v.Path)
							if err != nil {
								log.New(os.Stdout, "[Warning] graphql.rootQuery.posts",
									log.LstdFlags).Println(err.Error())
								return true
							}
							_v.Context = string(bytes)
							return false
						}
						_v = Post{}
						return true
					})
					if _v.Alias != "" {
						return _v, nil
					}
				}
				return nil, nil
			},
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(postType),
			Args: graphql.FieldConfigArgument{
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_list := []Post{}

				tagQ, ok := params.Args["tag"].(string)
				if ok {
					v, ok := TabMap.Load(tagQ)
					if ok {
						_v := v.([]Post)
						for i := 0; i < len(_v); i++ {
							_list = append(_list, _v[i])
						}
					}
					return _list, nil
				}

				return []Post{}, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: Query,
})

func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
