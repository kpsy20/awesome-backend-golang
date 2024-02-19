package sqlmap

import (
	"embed"
	"github.com/midir99/sqload"
)

//go:embed *.sql
var efs embed.FS

type Map struct {
	GetRoomList  string `query:"GetRoomList"`
	RegisterRoom string `query:"RegisterRoom"`
	RemoveRoom   string `query:"RemoveRoom"`
	ModifyRoom   string `query:"ModifyRoom"`
}

func New() *Map { return sqload.MustLoadFromFS[Map](efs) }
