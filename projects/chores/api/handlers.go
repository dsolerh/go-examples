package api

import (
	"chores_app/data"
	"net/http"
)

// CHORES

func (routes *Routes) GetAllChores(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) CreateChore(w http.ResponseWriter, r *http.Request) {
	dto := new(CreateChoreDTO)
	if err := DecodeJSON(r.Body, dto); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrorDTO{Error: "error parsing the request body"})
		return
	}

	if err := dto.Details.Validate(); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrorDTO{Error: "invalid chore details", Details: err.Error()})
		return
	}

	if err := dto.Schedule.Validate(); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrorDTO{Error: "invalid chore schedule", Details: err.Error()})
		return
	}

	chore, err := routes.Repository.CreateChore(data.ChoreDetails(*dto.Details), data.Schedule(*dto.Schedule))
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrorDTO{Error: "could not create the chore"})
		return
	}

	// return the chore
	RespondJSON(w, http.StatusCreated, transformChore(chore))
}

func (routes *Routes) GetChoreByID(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) UpdateChoreDetails(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) UpdateChoreSchedule(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) DeactivateChore(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

// MEMBERS

func (routes *Routes) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) CreateMember(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) GetMemberByID(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) UpdateMemberInfo(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}

func (routes *Routes) DeactivateMember(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}
