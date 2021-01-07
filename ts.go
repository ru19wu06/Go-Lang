package main

import(
	"fmt"
	"time"
)

func main(){
	arrA := [2]int{1,2}
	fmt.Printf("arrA : %p , %v\n", &arrA, arrA)

	arr(&arrA)
	arrA[0]++
	fmt.Printf("arrA : %p,%v\n",&arrA,arrA)

}

func arr(x*[2]int){
	fmt.Printf("pass Array : %p, %v\n",x,*x)
	time.Sleep(time.Second)
	(*x)[0]++
}


