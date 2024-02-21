package rooms

type RoomService interface {
	GetRoomList()
	RegisterRoom()
	ModifyRoom()
	RemoveRoom()
}
