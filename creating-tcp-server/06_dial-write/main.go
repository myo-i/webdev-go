package main

// 例えば02を起動している状態でこのコードを実行すると
// このコードは実行終了し、サーバー側でFprintlnで指定した文字列が出力される。

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// コネクションを得る
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// コネクションを得てエラーが発生しなかったら、下記を実行しdefer
	fmt.Fprintln(conn, "Dialed")
}
