package main

import "net"

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {

	//创建客户对象

	//创建Server

}

func main() {

}
