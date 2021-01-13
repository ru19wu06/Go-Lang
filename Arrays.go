package main

//結構範例
//費氏數列

import (
    "fmt"
)


func main(){
	scores := make([]int, 20)
	//scores[7] = 669
	scores[0]=1
	scores[1]=1
	for i:=0;i<17;i++{
		scores[i+2]=scores[i]+scores[i+1]

		fmt.Println(scores[i])
	}

	// scores = append(scores, 5)
	// scores = append(scores, 9)
	fmt.Println(scores) // prints [5]


}