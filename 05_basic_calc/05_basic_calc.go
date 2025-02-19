// # No. 05 四則演算
// 整数値を2つ入力させ、それらの値の和、差、積、商と余りを求めるプログラムを作成せよ。なお、差と商は1つ目の値から2つ目の値を引いた、あるいは割った結果とする。余りは無い場合も0と表示するのでよい。
package main

import "fmt"

func input_number(prompt string) (number int) {
	var num int
	fmt.Print(prompt)
	fmt.Scanln(&num)
	return num
}

func main() {
	num1 := input_number("input 1st number: ")
	num2 := input_number("input 2nd number: ")
	fmt.Println("和:", num1+num2)
	fmt.Println("差:", num1-num2)
	fmt.Println("積:", num1*num2)
	fmt.Println("商:", num1/num2, "余り:", num1%num2)
}
