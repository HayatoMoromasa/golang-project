package main

import "fmt"

//packageの設定は一つだけ

//変数

func main() {
	//明示的な定義
	//var 変数名　型　= 値
	var i int = 100
	fmt.Println(i)

	var s string = "hello go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "golang"
	)
	fmt.Println(i2, s2)
}
