package model

import (
	"fmt"
	"strings"

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
				aliasQ, ok := params.Args["alias"].(string)
				if ok {
					_post := Post_.posts
					for _, v := range _post {
						if v.Alias == aliasQ {
							return v, nil
						}
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
					_posts := Post_.posts
					_h := 0
					for i := offset; i < func() int {
						if offset > 0 {
							return len(_posts) - offset + 1
						}
						return len(_posts) - offset
					}(); i++ {
						if _h >= row && row != 0 {
							break
						}
						// 处理Content
						var _post = _posts[i]
						_post.Content = Content2Description(_post.Content, _post.Alias)

						if tagQ == "" {
							_list = append(_list, _post)
							_h++
						} else {
							for _, tag := range _post.Tags {
								if tagQ == tag {
									_list = append(_list, _post)
									_h++
								}
							}
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

func Content2Description(content, alias string) string {
	des := ""
	sp := strings.Split(content, "\n")
	for i := 0; i < len(sp); i++ {
		if strings.Index(sp[i], "```") == -1 &&
			strings.Index(sp[i], `<!-- More -->`) == -1 {
			des += sp[i] + "\n"
		} else {
			des += `<!-- More -->`
			break
		}
	}
	return des
}
