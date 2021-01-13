package main

import(
    "fmt"
)

func main(){
	lookup :=make(map[string]int)
	lookup["goku"] = 9001
	lookup["KUQU"] = 777
	lookup["vegeta"] = 3310
	power:= lookup["vegeta"]

	fmt.Println(power)
	fmt.Println(lookup["goku"])
	fmt.Println(lookup["KUQU"])
}