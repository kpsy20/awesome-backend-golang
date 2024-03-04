package repository

import (
	"awesome-backend-golang/internal/aggregate/rooms"
	"database/sql"
	"log/slog"
)

func (m *Repository) SelectRoomList() (res rooms.Rooms, err error) {
	var rows *sql.Rows
	if rows, err = m.db.Query(m.sql.GetRoomList); err != nil {
		return
	}

	for rows.Next() {
		var room rooms.Room
		err = rows.Scan(&room.Id, &room.Name, &room.AdminUserId, &room.CreateTime)
		if err != nil {
			return
		}
		res.Rooms = append(res.Rooms, room)
	}
	res.Success = true
	slog.Info("get payload list", "query", m.sql.GetRoomList, "list", res)
	return
}
