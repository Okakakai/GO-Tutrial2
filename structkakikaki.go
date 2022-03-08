package main

import "fmt"

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func main() {
	t := triangle{3}
	fmt.Println("Perimeter:", t.perimeter())
}

// type 構造体 struct {
//     // 定義
// }

// func (変数名 構造体) 関数名(変数名 型) 戻り値の型 {
//     // 処理
//     return
// }

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// // PersonのNameに文字列を結合する
// func (h Person) addName(f string) string {
// 	return h.Name + f
// }

// func main() {

// 	p := Person{
// 		Name: "me",
// 		Age:  20,
// 	}

// 	p.Name = p.addName("bee")

// 	fmt.Println(p)
// 	// {mebee 20}

// }
// javaのインスタンスみたいな使い方をするのね知らなかったわ
