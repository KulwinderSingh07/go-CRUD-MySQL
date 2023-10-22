package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KulwinderSingh07/go-bookstore/pkg/model"
	"github.com/KulwinderSingh07/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBook(res http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()
	book, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(book)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := model.GetBookById(Id)
	bookdata, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bookdata)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	createbook := &model.Book{}
	utils.ParseBody(req, createbook)
	// fmt.Println(req.Body)
	b := createbook.CreateBook()
	data, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := model.DeleteBook(Id)
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updateBook = &model.Book{}
	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while patrrsing")
	}
	bookDetails, db := model.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Name != "" {

		bookDetails.Name = updateBook.Name

	}
	if updateBook.Author != "" {

		bookDetails.Author = updateBook.Author

	}
	if updateBook.Publications != "" {

		bookDetails.Publications = updateBook.Publications

	}
	db.Save(&bookDetails)
	data, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}
