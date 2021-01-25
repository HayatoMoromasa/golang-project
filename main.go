package main

//packageの設定は一つだけ

import (
	"fmt"
	"time"
)

//↑複数importできる

/*
hello world
実行はgo run main.go
*/

func main() {
	fmt.Println("Hello world!!")
	fmt.Println(time.Now())
}
