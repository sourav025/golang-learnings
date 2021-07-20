package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func (book Book) update(oldAuthor *Book) {
	oldAuthor.Isbn = book.Isbn
	oldAuthor.Title = book.Title
	book.Author.update(oldAuthor.Author)
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (author Author) update(oldAuthor *Author) {
	oldAuthor.FirstName = author.FirstName
	oldAuthor.LastName = author.LastName
}

type Response struct {
    StatusCode int `json:"statusCode"`
    Msg string `json:"message"`
}


var books []Book

// Handler functions
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}



func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(&Response{404, "Item not found"})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	var found = false
	for index, item := range books {
		if item.ID == params["id"] {
			found = true
			newBook.ID = item.ID
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	if !found {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&Response{404, "Item not found"})
		return 
	}
	books = append(books, newBook)
	json.NewEncoder(w).Encode(&newBook)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false
	for index, item := range books {
		if item.ID == params["id"] {
			found = true
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	if !found {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&Response{404, "Item not found"})
		return
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{"1", "123123", "Book one", &Author{"Author1", "Author2"}})

	// Router handler
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
