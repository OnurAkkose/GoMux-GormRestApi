package api

import (
	"encoding/json"
	response "gorestapi/pkg/api"
	model "gorestapi/pkg/model/member"
	member "gorestapi/pkg/service/member"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MemberAPI struct {
	MemberService member.MemberService
}

func NewMemberAPI(a member.MemberService) MemberAPI {
	return MemberAPI{MemberService: a}
}

func (a MemberAPI) FindAllMembers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		books, err := a.MemberService.All()
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, books)
	}
}
func (a MemberAPI) FindById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		member, err := a.MemberService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToMemberDto(member))
	}
}

func (a MemberAPI) CreateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var authorDto model.MemberDto

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&authorDto); err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		defer r.Body.Close()

		createdMember, err := a.MemberService.Save(model.ToMember(authorDto))
		if err != nil {
			response.RespondWithJSON(rw, http.StatusOK, model.ToMemberDto(createdMember))
		}

	}
}
func (m MemberAPI) UpdateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		var memberDto model.MemberDto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&memberDto); err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		memberobj, err := m.MemberService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
		}

		memberobj.FullName = memberDto.FullName
		memberobj.Email = memberDto.Email
		updatedMember, err := m.MemberService.Save(memberobj)
		if err != nil {
			response.ResponseWithError(rw, http.StatusInternalServerError, err.Error())
		}
		response.RespondWithJSON(rw, http.StatusOK, model.ToMemberDto(updatedMember))
	}
}

func (m MemberAPI) DeleteMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response.ResponseWithError(rw, http.StatusBadRequest, err.Error())
			return
		}
		memberobj, err := m.MemberService.FindById(uint(id))
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		err = m.MemberService.Delete(memberobj.ID)
		if err != nil {
			response.ResponseWithError(rw, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		responseMessage := Response{
			Message: "Member deleted successfully!",
		}
		response.RespondWithJSON(rw, http.StatusOK, responseMessage)

	}
}

func (a MemberAPI) Migrate() {
	err := a.MemberService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
