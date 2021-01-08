package main

import (
    "fmt"
)

func main(){
    numbers := [...]int{0,1,2,3,4,5,6}
    fmt.Println(numbers)

    index :=3

	fmt.Println(numbers[index])
    fmt.Println(numbers[:index],numbers[index+1:])
	
	
	//去掉數字並合併
    deletedNumbers := append(numbers[:index],numbers[index+1:]...)
    fmt.Println(deletedNumbers)

}