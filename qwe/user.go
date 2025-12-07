package main

import "net"

type User struct {
	Name  string
	Adder string
	C     chan string
	Conn  net.Conn
}

func NewUser(conn net.Conn) *User {
	userAdder := conn.RemoteAddr().String()

	user := &User{
		Name:  userAdder,
		Adder: userAdder,
		C:     make(chan string),
		Conn:  conn,
	}
}
