package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iswarmondal/REST_Go/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("Endpoint hit: Home Page")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(models.Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	models.Articles = []models.Article{
		{
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		{
			Title:   "Hello 2",
			Desc:    "Article Description",
			Content: "Article Content",
		},
	}
	defer handleRequests()
}
