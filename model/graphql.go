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
			Type: postsType,
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
				_posts := Post_.posts

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
					_h := 0
					for i := (offset); i < len(_posts); i++ {
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
				}
				return struct {
					Max  int    `json:"max"`
					List []Post `json:"list"`
				}{len(_posts), _list}, nil
			},
		},
		"tags": &graphql.Field{
			Type: graphql.NewList(tagType),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				type t struct {
					Tag   string `json:"tag"`
					Posts []Post `json:"posts"`
				}
				var tags []t
				TagMap.Range(func(k, v interface{}) bool {
					_key := k.(string)
					tags = append(tags, t{
						Tag:   _key,
						Posts: v.([]Post),
					})
					return true
				})
				// log.Println(tags)
				return tags, nil
			},
		},
		"links": &graphql.Field{
			Type: graphql.NewList(linkType),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return Link_.links, nil
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
