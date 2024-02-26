-- query: GetRoomList
SELECT id, name, admin_user_id, create_time FROM room;

-- query: RegisterRoom
INSERT INTO room (name, id, admin_user_id) VALUES (?, ?, ?);

-- query: RemoveRoom
DELETE FROM room WHERE id = ?;

-- query: ModifyRoom
UPDATE room SET name = ? WHERE id = ?;