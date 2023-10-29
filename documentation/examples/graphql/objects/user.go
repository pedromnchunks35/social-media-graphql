package objects

import "github.com/graphql-go/graphql"

/*
? Post Type, this will contain the normal id,title and description
*/
var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

/*
? User Type, this will contain the id, name and posts
*/
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"posts": &graphql.Field{
				Type: graphql.NewList(PostType),
			},
		},
	},
)
