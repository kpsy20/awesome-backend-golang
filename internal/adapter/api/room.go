package api

import (
	"awesome-backend-golang/internal/aggregate/rooms"
	"github.com/go-chi/render"
	"net/http"
)

func (a *API) GetRoomList(res http.ResponseWriter, req *http.Request) {
	reqBody := &rooms.Room{}
	if err := render.Decode(req, reqBody); err != nil {
		HandleError(res, req, err)
		return
	}

	roomList, err := a.rooms.GetRoomList()
	if err != nil {
		HandleError(res, req, err)
		return
	}

	render.Respond(res, req, roomList)
}

func (a *API) RegisterRoom(res http.ResponseWriter, req *http.Request) {

}
func (a *API) ModifyRoom(res http.ResponseWriter, req *http.Request) {

}
func (a *API) RemoveRoom(res http.ResponseWriter, req *http.Request) {

}
