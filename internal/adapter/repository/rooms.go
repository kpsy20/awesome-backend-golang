package repository

import (
	"awesome-backend-golang/internal/aggregate/rooms"
	"log/slog"
)

func (m *Repository) SelectRoomList() (rooms.Rooms, error) {
	slog.Info("select room list")
	var res rooms.Rooms
	rows, err := m.db.Query(m.sql.SelectRoomList)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var room rooms.Room
		err := rows.Scan(&room.Id, &room.Name, &room.AdminUserId, &room.CreateTime)
		if err != nil {
			return res, err
		}
		res.Rooms = append(res.Rooms, room)
	}
	return res, err
}

func (m *Repository) SelectRoom(name string, adminUserId int) (rooms.Rooms, error) {
	slog.Info("select room", "query", m.sql.SelectRoomByNameAndAdminUserId, "name", name, "admin_user_id", adminUserId)
	var res rooms.Rooms
	rows, err := m.db.Query(m.sql.SelectRoomByNameAndAdminUserId, name, adminUserId)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var room rooms.Room
		err := rows.Scan(&room.Id, &room.Name, &room.AdminUserId, &room.CreateTime)
		if err != nil {
			return res, err
		}
		res.Rooms = append(res.Rooms, room)
	}
	return res, err
}

func (m *Repository) SelectRoomById(id int) (rooms.Rooms, error) {
	slog.Info("select room", "query", m.sql.SelectRoomById, "id", id)
	var res rooms.Rooms
	rows, err := m.db.Query(m.sql.SelectRoomById, id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var room rooms.Room
		err := rows.Scan(&room.Id, &room.Name, &room.AdminUserId, &room.CreateTime)
		if err != nil {
			return res, err
		}
		res.Rooms = append(res.Rooms, room)
	}
	return res, err
}

func (m *Repository) InsertRoom(name string, adminUserId int) error {
	_, err := m.db.Exec(m.sql.InsertRoom, name, adminUserId)
	if err != nil {
		return err
	}
	return err
}

func (m *Repository) UpdateRoom(id int, newName string) error {
	_, err := m.db.Exec(m.sql.UpdateRoom, newName, id)
	if err != nil {
		return err
	}
	return err
}

func (m *Repository) DeleteRoom(id int) error {
	_, err := m.db.Exec(m.sql.DeleteRoom, id)
	if err != nil {
		return err
	}
	return err
}
