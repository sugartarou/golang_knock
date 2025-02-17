// # No. 04 入力と計算
// 整数値を入力させ、その入力値を3倍した計算結果を表示するプログラムを作成せよ。
package main

import "fmt"

func main() {
	var number int
	fmt.Print("input number: ")
	fmt.Scanln(&number)
	fmt.Println("answer =", number*3)
}
