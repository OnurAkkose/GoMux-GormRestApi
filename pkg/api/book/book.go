package api

import (
	"encoding/json"
	response "gorestapi/pkg/api"
	model "gorestapi/pkg/model/book"
	book "gorestapi/pkg/service/book"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookAPI struct {
	BookService book.BookService
}

func NewBookAPI(a book.BookService) BookAPI {
	return BookAPI{BookService: a}
}

func (a BookAPI) FindAllBooks() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		books, err := a.BookService.All()
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, books)
	}
}
func (a BookAPI) FindById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		book, err := a.BookService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToBookDto(book))
	}
}

func (a BookAPI) CreateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var authorDto model.BookDto

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&authorDto); err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		defer r.Body.Close()

		createdAuthor, err := a.BookService.Save(model.ToBook(authorDto))
		if err != nil {
			response.RespondWithJSON(rw, http.StatusOK, model.ToBookDto(createdAuthor))
		}

	}
}

func (b BookAPI) UpdateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		var bookDto model.BookDto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&bookDto); err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		book, err := b.BookService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
		}

		book.Title = bookDto.Title
		book.AuthorId = bookDto.AuthorId
		updatedBook, err := b.BookService.Save(book)
		if err != nil {
			response.ResponseWithError(rw, http.StatusInternalServerError, err.Error())
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToBookDto(updatedBook))
	}
}

func (b BookAPI) DeleteBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}
		book, err := b.BookService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		err = b.BookService.Delete(book.ID)
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		responseMessage := Response{
			Message: "Book deleted successfully!",
		}
		response.RespondWithJSON(rw, http.StatusOK, responseMessage)

	}
}

func (a BookAPI) Migrate() {
	err := a.BookService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
