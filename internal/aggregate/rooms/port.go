package rooms

type RoomService interface {
	GetRoomList() (roomList Rooms, err error)
	RegisterRoom() (roomList Rooms, err error)
	ModifyRoom() (roomList Rooms, err error)
	RemoveRoom() (err error)
}

type Repository interface {
	SelectRoomList() (res Rooms, err error)
}
