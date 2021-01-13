package main

import(
    "fmt"
    "math/rand"
    "sort"
)

func main(){
    scores := make([]int,100)
	
	for u:=0;u<10;u++{
		scores[u] = int(rand.Int31n(1000))
		fmt.Println(scores[u])
	}

	sort.Ints(scores)

	worst := make([]int,5)

	copy(worst,scores[:5])
	fmt.Println(worst)
	for s:=0;s<2;s++{
		for i:=0;i<11;i++{
			fmt.Println(rand.Int31n(1000))
			
		}
		fmt.Println("----------------")
	}

}