package rooms

import (
	"errors"
)

type rooms struct {
	repo Repository
}

func New(repo Repository) RoomService {
	return &rooms{repo: repo}
}

func (r *rooms) GetRoomList() (Rooms, error) {
	roomList, err := r.repo.SelectRoomList()
	return roomList, err
}

type (
	RegisterRoomReq struct {
		Name        string `json:"name" xml:"name" db:"name"`
		AdminUserId int    `json:"admin_user_id" xml:"admin_user_id" db:"admin_user_id"`
	}
)

func (r *rooms) RegisterRoom(req RegisterRoomReq) (Rooms, error) {
	roomList, err := r.repo.SelectRoom(req.Name, req.AdminUserId)
	if err != nil {
		return roomList, err
	}

	if len(roomList.Rooms) != 0 {
		return roomList, errors.New("room name is duplicated")
	}

	err = r.repo.InsertRoom(req.Name, req.AdminUserId)
	if err != nil {
		return roomList, err
	}

	roomList, err = r.repo.SelectRoom(req.Name, req.AdminUserId)
	if err != nil {
		return roomList, err
	}
	return roomList, err
}

type (
	ModifyRoomReq struct {
		Id      int    `json:"id" xml:"id" db:"id"`
		NewName string `json:"new_name" xml:"new_name" db:"new_name"`
	}
)

func (r *rooms) ModifyRoom(req ModifyRoomReq) (Rooms, error) {
	err := r.repo.UpdateRoom(req.Id, req.NewName)
	if err != nil {
		return Rooms{}, err
	}

	roomList, err := r.repo.SelectRoomById(req.Id)
	if err != nil {
		return roomList, err
	}
	return roomList, err
}

type (
	RemoveRoomReq struct {
		Id int `json:"id" xml:"id" db:"id"`
	}
)

func (r *rooms) RemoveRoom(req RemoveRoomReq) error {
	err := r.repo.DeleteRoom(req.Id)
	if err != nil {
		return err
	}
	return err
}
