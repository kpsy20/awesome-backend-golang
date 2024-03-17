-- query: SelectRoomList
select id, name, admin_user_id, create_time from room;

-- query: SelectRoomByNameAndAdminUserId
select id, name, admin_user_id, create_time from room where name = ? and admin_user_id = ?;

-- query: SelectRoomById
select id, name, admin_user_id, create_time from room where id = ?;

-- query: InsertRoom
insert into room (name, admin_user_id) values (?, ?);

-- query: UpdateRoom
update room set name = ? where id = ?;

-- query: DeleteRoom
delete from room where id = ?;