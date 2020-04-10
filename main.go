package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type post struct {
	Title  string
	Points int32
}

var posts = []post{
	{Title: "TITLE_1"},
	{Title: "TITLE_2"},
}

var queryObjectConfig = graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(postObjectConfig)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return posts, nil
			},
		},
	},
}

var postObjectConfig = graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"points": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "POINTS HERE", nil
			},
		},
	},
}

var queryType = graphql.NewObject(queryObjectConfig)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", h)

	log.Print("Starting listener at port ", port)
	http.ListenAndServe(":"+port, nil)

}
