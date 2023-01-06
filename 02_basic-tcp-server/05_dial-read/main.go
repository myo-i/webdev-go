package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// コネクション取得
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// コネクションを読み取る
	go read(conn)
	// bs, err := ioutil.ReadAll(conn)
	data := make([]byte, 1024)
	// fmt.Println(string(bs))

	count, err := conn.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data[:count]))

}

func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println("------------")
}
