package main

import (
	"fmt"
	"sort"
)

func GuessingGame() {
	var num int
	fmt.Println("输入一个0-100的整数")
	fmt.Scanf("%d", &num)
	answer := sort.Search(100, func(i int) bool {
		if i < num {
			fmt.Printf("你输入的数字 > %d \n", i)
			return false
		} else {
			fmt.Printf("你输入的数字 <= %d \n", i)
		}
		return true
	})
	fmt.Printf("你输入的数字是 %d \n", answer)
}

func main() {
	GuessingGame()
}
