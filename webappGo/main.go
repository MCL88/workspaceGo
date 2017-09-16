package main

import
(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct{
	Title 		string 	`json:"Title"`
	Desc 		string	`json:"desc"`	
	Content 	string	`json:"content"`	
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Ciao", Desc: "Descrizione 1", Content:"Ciao mondo!"},
		Article{Title:"Ciao2", Desc: "Descrizione 2", Content:"Loda il sole, stolto"},
	}
	fmt.Println("Fine returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Benvenuto nella Home!")
	fmt.Println("Fine homePage")
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	HandleRequest()
}