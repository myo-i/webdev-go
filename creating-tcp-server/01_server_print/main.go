package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listenの定義を見るとListenerというインターフェースを継承していることがわかる。
	// そしてListenerはAccept,Close, Addrというメソッドを定義していることから
	// Listenの返り値であるnet.Listener型はCloseしなければいけないということがわかる。
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	var index int
	for {
		// コネクション毎にfor文が回っているのか見てみる
		index++
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		if index == 1 {
			io.WriteString(connection, "\nWelcome My TCP Server!!\n")
			fmt.Fprintln(connection, "Hellooooo")
			fmt.Fprintf(connection, "%v", "Good luck!!\n")
		} else {
			fmt.Fprintln(connection, "Connection is over one times")
			fmt.Println(index, ": times accessed!!")
		}

		connection.Close()
	}
}
