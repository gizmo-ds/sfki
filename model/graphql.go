package model

import (
	"fmt"

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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_v := Post{}

				aliasQ, ok := params.Args["alias"].(string)
				if ok {
					Posts.Range(func(_, v interface{}) bool {
						_v = v.(Post)
						if _v.Alias == aliasQ {
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
				"offset": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"row": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_list := []Post{}
				offset, ok := params.Args["offset"].(int)
				if !ok || offset < 0 {
					offset = 0
				}
				row, ok := params.Args["row"].(int)
				if !ok || row < 0 {
					row = 0
				}

				tagQ, ok := params.Args["tag"].(string)
				if ok {
					v, ok := TabMap.Load(tagQ)
					if ok {
						_v := v.([]Post)
						_h := 0
						for i := offset; i < func() int {
							if offset > 0 {
								return len(_v) - offset + 1
							}
							return len(_v) - offset
						}(); i++ {
							if _h >= row && row != 0 {
								break
							}
							_list = append(_list, _v[i])
							_h++
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
