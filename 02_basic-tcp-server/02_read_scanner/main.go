package main

// サーバーを立てたら localhost:8080やlocalhosst:8080/Golang-developmentなどを打っみると
// GET / や GET /Golang-development がコンソールに表示される

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// 基本的にはここには辿り着かない
	// ストリーム接続が開かれている
	// 上記のリーダーはどうやって完了を知るのか？
	fmt.Println("Code got here.")
}
