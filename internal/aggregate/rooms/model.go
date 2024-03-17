package rooms

type (
	Room struct {
		Name        string `json:"name" xml:"name" db:"name"`
		Id          int    `json:"id" xml:"id" db:"id"`
		AdminUserId int    `json:"admin_user_id" xml:"admin_user_id" db:"admin_user_id"`
		CreateTime  string `json:"create_time" xml:"create_time" db:"create_time"`
	}
	Rooms struct {
		Rooms []Room `json:"rooms" xml:"rooms"`
	}
)
