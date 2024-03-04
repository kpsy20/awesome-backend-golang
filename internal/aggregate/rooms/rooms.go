package rooms

type rooms struct {
	repo Repository
}

func New(repo Repository) RoomService {
	return &rooms{repo: repo}
}

func (r *rooms) GetRoomList() (roomList Rooms, err error) {
	roomList, err = r.repo.SelectRoomList()
	return
}

func (r *rooms) RegisterRoom() (roomList Rooms, err error) {
	return
}

func (r *rooms) ModifyRoom() (roomList Rooms, err error) {
	return
}
func (r *rooms) RemoveRoom() (err error) {
	return
}
