package main

import (
	"fmt"
	"net"
)

// 设置服务端口
type Server struct {
	Ip   string
	Port int
}

// 返回新连接
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
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
	fmt.Println("建立链接")
}
