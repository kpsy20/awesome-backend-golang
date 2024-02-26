package rooms

type RoomService interface {
	GetRoomList() (roomList *Rooms, err error)
	RegisterRoom()
	ModifyRoom()
	RemoveRoom()
}

type Repository interface {
	SelectRoomList() (res *Rooms, err error)
}
