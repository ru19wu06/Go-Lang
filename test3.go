package main

import (
    "fmt"
)

func main(){
	
	numbers := make([]int,0)

	for i :=0 ;i<20;i++ {
		numbers = append(numbers,i)
		fmt.Printf("len: %d,cap: %d,ptr: %p\n",len(numbers),cap(numbers),numbers)
	}

	
}