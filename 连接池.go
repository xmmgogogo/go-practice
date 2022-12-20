package main

import (
	"fmt"
)

type Client struct {
	Name string
}

type MysqlConn struct {
	Conn []Client
	Num  int
}

func NewConn() *MysqlConn {
	return &MysqlConn{}
}

func (t *MysqlConn) AddOneClient(name string) {
	t.Conn = append(t.Conn, Client{Name: name})
}

func (t *MysqlConn) GetOneClient() Client {
	// 随机算法
	//rand.Seed(time.Now().UnixNano())
	//r := rand.Intn(len(t.Conn))
	//
	//return t.Conn[r]

	// 轮询算法
	randClient := t.Conn[t.Num%len(t.Conn)] // 1/3
	t.Num++
	if t.Num >= len(t.Conn) {
		t.Num = 0
	}
	return randClient
}

func main() {
	g := NewConn()
	g.AddOneClient("cat")
	g.AddOneClient("dog")
	g.AddOneClient("fish")

	fmt.Println("conn list", g.Conn)

	fmt.Println("返回的链接：", g.Num, g.GetOneClient())
	fmt.Println("返回的链接：", g.Num, g.GetOneClient())
	fmt.Println("返回的链接：", g.Num, g.GetOneClient())
	fmt.Println("返回的链接：", g.Num, g.GetOneClient())
	fmt.Println("返回的链接：", g.Num, g.GetOneClient())
}
