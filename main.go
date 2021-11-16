package main

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}

var books []Book

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: welcome")
	fmt.Fprintf(w, "welcome!")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: addBook")

	// read POST request
	body, _ := ioutil.ReadAll(r.Body)

	var newBook Book
	json.Unmarshal([]byte(body), &newBook)

	books = append(books, newBook)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getAllBooks")
	json.NewEncoder(w).Encode(books)
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/add", addBook)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	books = []Book{
		Book{Title: "Stalingrad", Author: "Antony Beevor", ISBN: "9780141032405"},
		Book{Title: "Symphony for the City of the Dead", Author: "M. T. Anderson", ISBN: "97807763691004"},
		Book{Title: "A People's History of the United States", Author: "Howard Zinn", ISBN: "9780062397348"},
	}

	handleRequests()
}
