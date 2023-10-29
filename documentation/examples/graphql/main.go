package main

import (
	"encoding/json"
	"fmt"
	"getting-started/database"
	"getting-started/execution"
	"getting-started/schema"
	"getting-started/templates"
	"log"
	"net/http"
)

func main() {
	//? init data
	data := &map[int]*templates.User{}
	database.Data = data
	(*data)[1] = &templates.User{
		ID:   1,
		Name: "Pedro",
		Posts: []templates.Post{
			templates.Post{
				ID:          1,
				Title:       "O horrivel aconteceu",
				Description: "Aconteceu alguma coisa",
			},
		},
	}
	//? Create a schema
	sch, err := schema.CreateSchema()
	if err != nil {
		log.Fatalf("Something went wrong creating the schema: %v", err)
	}
	//? Make it available in http query
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := execution.ExecuteQuery(r.URL.Query().Get("query"), sch)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
