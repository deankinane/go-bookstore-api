package controllers

import (
	"net/http"

	"github.com/deankinane/go-bookstore-api/pkg/models"
	"github.com/deankinane/go-bookstore-api/pkg/utils"
)

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if success := utils.ParseBody(w, r, book); !success {
		return
	}

	utils.Json201Response(w, book.CreateBook())

}

var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	utils.Json200Response(w, books)
}

var GetBookById = func(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdParam(w, r)
	if id == 0 {
		return
	}

	book := models.GetBookById(id)
	if book == nil {
		utils.Json404Response(w)
		return
	}

	utils.Json200Response(w, book)
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdParam(w, r)
	if id == 0 {
		return
	}

	book := &models.Book{}
	if success := utils.ParseBody(w, r, book); !success {
		return
	}

	utils.Json200Response(w, models.UpdateBook(id, book))
}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdParam(w, r)
	if id == 0 {
		return
	}

	book := models.DeleteBook(id)
	if book == nil {
		utils.Json404Response(w)
		return
	}

	utils.Json200Response(w, book)
}
