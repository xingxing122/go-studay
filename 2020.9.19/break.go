package main

import "fmt"

func main() {
	for  i  := 0; i < 10 ; i++{
		if i ==5{
			continue // 跳过本次循环，继续下一次
			//break  // 退出循环
		}
		fmt.Println(i)
	}

}
