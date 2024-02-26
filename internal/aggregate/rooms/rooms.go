package rooms

type rooms struct {
	repo Repository
}

func New(repo Repository) RoomService {
	return &rooms{repo: repo}
}

func (r *rooms) GetRoomList() (roomList *Rooms, err error) {
	roomList, err = r.repo.SelectRoomList()
	return
}

func (r *rooms) RegisterRoom() {

}

func (r *rooms) ModifyRoom() {

}
func (r *rooms) RemoveRoom() {

}
