package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

//创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线的用户
func (this *Server) ListenMessager()  {
	for  {
		msg := <-this.Message
		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _,cli := range this.OnlineMap{
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}


//广播地址的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + "[" + user.Name + "]" + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handle(conn net.Conn) {
	//当前链接的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn)

	//用户上线 加到onlinemap 中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播当前用户上线消息
	this.BroadCast(user,"shang xian")

	//接受用户发送来的消息
	go func(){
		buf :=  make([]byte,4096)

		for  {
			n,err := conn.Read(buf)
			if n == 0 {
				this.BroadCast(user,"xia xian")
				return
			}

			if err != nil && err != io.EOF{
				fmt.Println("Conn Read err:",err)
				return
			}

			//提取用户的消息 去掉“\n”
			msg := string(buf[:n-1])

			//消息广播
			this.BroadCast(user,msg)
		}

	}()


	//当前handle阻塞
	select {}

}

//启动服务器的方法
func (this *Server) Start() {

	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}

	//close
	defer listener.Close()

	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listenner accept err:", err)
			continue
		}

		//do handle
		go this.Handle(conn)
	}

	//do handle

}