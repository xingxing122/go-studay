package main

import "fmt"
func main() {
	num  := []int{1,2,3,4,5,6,}
	fmt.Println(len(num),cap(num))

	num = append(num,6)
	fmt.Println(len(num),cap(num))

}
