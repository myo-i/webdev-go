package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var writeDatas []string

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
		handle(conn)
		for _, data := range writeDatas {
			fmt.Println(data)
			byteData := []byte(data)
			_, err := conn.Write(byteData)
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		writeDatas = append(writeDatas, ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}
