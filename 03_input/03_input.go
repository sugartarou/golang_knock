// # No. 03 入力
// 整数値を入力させ、その入力値を表示するプログラムを作成せよ。
package main

import "fmt"

func main() {
	var number int
	fmt.Print("input number: ")
	fmt.Scanln(&number)
	fmt.Println("your number is ", number)
}
