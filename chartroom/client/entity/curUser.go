package entity

import (
	"gocode/chartroom/server/model"
	"net"
)

type CurUser struct {
	Conn net.Conn
	model.User
}
