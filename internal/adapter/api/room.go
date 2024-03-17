package api

import (
	"awesome-backend-golang/internal/aggregate/rooms"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"regexp"
)

var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9가-힣!@#$%^&*()_+=-][a-zA-Z0-9가-힣!@#$%^&*()_+=-\\\s]{0,18}[a-zA-Z0-9가-힣!@#$%^&*()_+=-]$`)

func (a *API) GetRoomList(res http.ResponseWriter, req *http.Request) {
	roomList, err := a.rooms.GetRoomList()
	if err != nil {
		HandleError(res, req, err)
		return
	}

	render.Respond(res, req, roomList)
}

type (
	RegisterRoomReq struct {
		Name        string `json:"name" xml:"name" db:"name"`
		AdminUserId int    `json:"admin_user_id" xml:"admin_user_id" db:"admin_user_id"`
	}
)

func (a *API) RegisterRoom(res http.ResponseWriter, req *http.Request) {
	reqBody := &RegisterRoomReq{}
	if err := render.Decode(req, reqBody); err != nil {
		HandleError(res, req, err)
		return
	}
	if !validNamePattern.MatchString(reqBody.Name) {
		HandleError(res, req, errors.New("the name does not follow the rules"))
		return
	}
	roomReq := rooms.RegisterRoomReq{Name: reqBody.Name, AdminUserId: reqBody.AdminUserId}
	roomList, err := a.rooms.RegisterRoom(roomReq)
	if err != nil {
		HandleError(res, req, err)
		return
	}

	render.Respond(res, req, roomList)
}

type (
	ModifyRoomReq struct {
		Id      int    `json:"id" xml:"id" db:"id"`
		NewName string `json:"new_name" xml:"new_name" db:"new_name"`
	}
)

func (a *API) ModifyRoom(res http.ResponseWriter, req *http.Request) {
	reqBody := &ModifyRoomReq{}

	if err := render.Decode(req, reqBody); err != nil {
		HandleError(res, req, err)
		return
	}

	if !validNamePattern.MatchString(reqBody.NewName) {
		HandleError(res, req, errors.New("the name does not follow the rules"))
		return
	}

	roomReq := rooms.ModifyRoomReq{Id: reqBody.Id, NewName: reqBody.NewName}
	roomList, err := a.rooms.ModifyRoom(roomReq)
	if err != nil {
		HandleError(res, req, err)
		return
	}

	render.Respond(res, req, roomList)
}

type (
	RemoveRoomReq struct {
		Id int `json:"id" xml:"id" db:"id"`
	}
)

func (a *API) RemoveRoom(res http.ResponseWriter, req *http.Request) {
	reqBody := &RemoveRoomReq{}
	if err := render.Decode(req, reqBody); err != nil {
		HandleError(res, req, err)
		return
	}

	roomReq := rooms.RemoveRoomReq{Id: reqBody.Id}
	err := a.rooms.RemoveRoom(roomReq)
	if err != nil {
		HandleError(res, req, err)
		return
	}

	render.Respond(res, req, rooms.Rooms{})
}
