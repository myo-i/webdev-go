package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Hoge Huga"
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>WebDev myo-i</title>
</head>
<body>
<h1>` +
		name +
		`</h1>
</body>
</html>
	`)

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(str))
}
