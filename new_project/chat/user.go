package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

//创建一个用的API
func NewUser(conn net.Conn, server *Server) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//启动监听当前user  channel 消息的goroutine
	go user.ListenMessage()

	return user
}

//监听当前User channel的方法 一旦有消息，就直接发送给对应的客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//用户上线业务
func (this *User) Online() {
	//用户上线 加到onlinemap 中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "shang xian ...")
}

//用户下线业务
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	this.server.BroadCast(this, "xia xian...")
}

//给对应用户发消息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//用户处理消息的业务
func (this *User) DnMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有那些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":online ing...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {

		newName := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMsg("this name is exited...")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.sendMsg("rename success:" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {

		//消息模式：to|张三|消息内容

		//1.获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.sendMsg("msg type is error")
		}

		//2根据用户名 得到对方的User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.sendMsg("user is not exite")
		}

		//3获取消息内容，通过对方User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.sendMsg("content is null ...")
			return
		}
		remoteUser.sendMsg(this.Name + " to you say:" + content + "\n")

	} else {
		this.server.BroadCast(this, msg)
	}
}
