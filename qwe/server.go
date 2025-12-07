package main

import (
	"fmt"
	"net"
	"sync"
)

// 设置服务端口
type Server struct {
	Ip   string
	Port int

	//
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

// 返回新连接
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 开始运行
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpet err", err)
			continue
		}
		go s.Handler(conn)
	}
}

// 进行事物处理
func (s *Server) Handler(con net.Conn) {
	//fmt.Println("建立链接")
	user := NewUser(con)
	//用户上线
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	//广播当前上线消息
	s.BoradCast(user, "已上线")
}

func (s *Server) BoradCast(user *User, msg string) {
	sendMsg := "[" + user.Adder + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}
