package api

import (
	"encoding/json"
	response "gorestapi/pkg/api"
	model "gorestapi/pkg/model/author"
	author "gorestapi/pkg/service/author"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AuthorAPI struct {
	AuthorService author.AuthorService
}

func NewAuthorAPI(a author.AuthorService) AuthorAPI {
	return AuthorAPI{AuthorService: a}
}

func (a AuthorAPI) FindAllAuthors() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		authors, err := a.AuthorService.All()
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, authors)
	}
}
func (a AuthorAPI) FindById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		author, err := a.AuthorService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToAuthorDto(author))
	}
}

func (a AuthorAPI) CreateAuthor() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var authorDto model.AuthorDto

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&authorDto); err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		defer r.Body.Close()

		createdAuthor, err := a.AuthorService.Save(model.ToAuthor(authorDto))
		if err != nil {
			response.RespondWithJSON(rw, http.StatusOK, model.ToAuthorDto(createdAuthor))
		}

	}
}

func (a AuthorAPI) UpdateAuthor() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		var authorDto model.AuthorDto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&authorDto); err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		author, err := a.AuthorService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
		}

		author.FullName = authorDto.FullName
		author.Email = authorDto.Email
		updatedAuthor, err := a.AuthorService.Save(author)
		if err != nil {
			response.ResponseWithError(rw, http.StatusInternalServerError, err.Error())
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToAuthorDto(updatedAuthor))
	}
}

func (a AuthorAPI) DeleteAuthor() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}
		author, err := a.AuthorService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		err = a.AuthorService.Delete(author.ID)
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		responseMessage := Response{
			Message: "Author deleted successfully!",
		}
		response.RespondWithJSON(rw, http.StatusOK, responseMessage)

	}
}

func (a AuthorAPI) Migrate() {
	err := a.AuthorService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
