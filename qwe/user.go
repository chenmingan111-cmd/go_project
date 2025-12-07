package main

import "net"

type User struct {
	Name  string
	Adder string
	C     chan string
	conn  net.Conn
}

func NewUser(conn net.Conn) *User {
	userAdder := conn.RemoteAddr().String()

	user := &User{
		Name:  userAdder,
		Adder: userAdder,
		C:     make(chan string),
		conn:  conn,
	}
	//启动监听user channel消息
	go user.listenMessage()

	return user
}

func (u *User) listenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
