package main

import(
	"fmt"
	"reflect"
)

func main(){
	var arr = [...]int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(arr[8])

	slice0 := arr[:5]
	fmt.Println(slice0)

	slice1 := arr[5:]
	fmt.Println(slice1)
	
	slice2 := arr[2:7]
	fmt.Println(slice2)

	slice3 := arr[:]
	fmt.Println(slice3)

	slice4 := arr[0:0]
	fmt.Println(slice4)
	fmt.Println("------------")

	slice := arr[1:3]
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(len(slice),cap(slice))

	fmt.Println(arr)
	fmt.Println(slice)

	fmt.Println("-------------")
	slice[0]=0
	fmt.Println(arr)
	fmt.Println(slice)
}
