package main

//陣列結構
//陣列拔出
import (
    "fmt"
)


func main() {
	scores := []int{1, 2, 3, 4, 5,7}
	fmt.Println(len(scores))
	scores = removeAtIndex(scores, 2)
	
	fmt.Println(scores)
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
	