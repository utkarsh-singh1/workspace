package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	bookapi "github.com/utkarsh-singh1/bookApI/bookAPI"
)

type bookServer struct {

	store *bookapi.BookDb

}

func new() *bookServer {

	store := bookapi.New()
	return &bookServer{store: store}

}

func (bs *bookServer) CreatetaskHandler(w http.ResponseWriter, req *http.Request) {

	type RequestTask struct {

		name string
		author string
		date time.Time
		
	}

	type ResponseTask struct {

		Id int
	}
	
	var rt RequestTask


	// Response from the server
	
	id := bs.store.CreateEntry(rt.name, rt.author, rt.date)

	js, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (bs *bookServer) GetTaskHandler(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	book, err := bs.store.GetById(id)
	if err != nil {
		http.Error(w,fmt.Sprintf("Sorry we couldn't find the book by the id - %v ",id) , http.StatusInternalServerError)
	}

	js, err := json.Marshal(book)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	

}

func (bs *bookServer) GetAllHandler(w http.ResponseWriter, req *http.Request ) {

	books := bs.store.GetAllEntry()

	js , err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(js)

}

func (bs *bookServer) GetByAuthorHandler(w http.ResponseWriter, req *http.Request) {

	author := req.PathValue("author")


	// Response body

	books := bs.store.GetByAuthor(author)

	js, err := json.Marshal(books)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
	

}

func (bs *bookServer) DeleteByIdHandler(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(req.PathValue("id")) 
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = bs.store.DeleteEntry(id)
	if err != nil {

		http.Error(w,fmt.Sprintf("Could not find the book by the ID - %v", id), http.StatusNotFound)
	}
	
}

func (bs *bookServer) DeleteAllHandler(w http.ResponseWriter, req *http.Request) {

	bs.store.DeleteAll()

}

func main() {

	newserver := new()
	mux := http.NewServeMux()

	mux.HandleFunc("/book/", newserver.CreatetaskHandler )
	mux.HandleFunc("/book/", newserver.GetAllHandler)
	mux.HandleFunc("/book/{id}", newserver.GetTaskHandler)
	mux.HandleFunc("/book/author", newserver.GetByAuthorHandler)
	mux.HandleFunc("/book/", newserver.DeleteAllHandler)
	mux.HandleFunc("/book/{id}", newserver.DeleteByIdHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
	
}
