package main

import "fmt"

func bubble (nums []int){
	for j :=0; j < len(nums)-1; j++ {
		for i := 0; i < len(nums)-1-j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

}


func   main(){
	nums := []int{100,3000,30,40000,5000,50}

	fmt.Println(bubble(nums))

}