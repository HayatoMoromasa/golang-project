package main

import "fmt"

//packageの設定は一つだけ

//変数

// i5 := 500
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

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

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 5
	s3 = "go"
	fmt.Println(i3, s3)
	//上書きできる
	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//変数 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 500
	fmt.Println(i4)
	fmt.Println(i5)
	// i4 := 500のように再定義はできない
	// i4 = "hey"のようにintで定義したのにstringで再定義できない
	//9行目i5のように関数外での定義もできない

	//関数の呼び出し
	outer()
	// fmt.Println(s4)これは呼べない(s4はouter()関数のなかだけで有効だから)

	// var s5 string = "s5"goでは定義された変数は必ず使わないといけない

}
