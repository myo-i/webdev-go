package main

// サーバーを立ててtelnet localhost 8080 でアクセス
// アクセスした方から文字列を打ち込む

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
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
		fmt.Println("----" + ln)
		fmt.Fprintf(conn, "Input: %s\n", ln)
	}
	defer conn.Close()

	// 基本的にはここには辿り着かない
	// ストリーム接続が開かれている
	// 上記のリーダーはどうやって完了を知るのか？
	fmt.Println("Code got here.")
}
