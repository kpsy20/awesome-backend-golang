package rooms

type RoomService interface {
	GetRoomList() (Rooms, error)
	RegisterRoom(RegisterRoomReq) (Rooms, error)
	ModifyRoom(ModifyRoomReq) (Rooms, error)
	RemoveRoom(RemoveRoomReq) error
}

type Repository interface {
	SelectRoomList() (Rooms, error)
	SelectRoom(string, int) (Rooms, error)
	SelectRoomById(int) (Rooms, error)
	InsertRoom(string, int) error
	UpdateRoom(int, string) error
	DeleteRoom(id int) error
}
