package sqlmap

import (
	"embed"
	"github.com/midir99/sqload"
)

//go:embed *.sql
var efs embed.FS

type Map struct {
	SelectRoomList                 string `query:"SelectRoomList"`
	SelectRoomByNameAndAdminUserId string `query:"SelectRoomByNameAndAdminUserId"`
	SelectRoomById                 string `query:"SelectRoomById"`
	InsertRoom                     string `query:"InsertRoom"`
	DeleteRoom                     string `query:"DeleteRoom"`
	UpdateRoom                     string `query:"UpdateRoom"`
}

func New() *Map { return sqload.MustLoadFromFS[Map](efs) }
