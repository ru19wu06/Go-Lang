package main

import(
    "fmt"
)

func main(){
    numbers := make([]int,0)

    for i:=0;i<10;i++{
        numbers = append(numbers,i)
    }

    copyA :=make([]int,len(numbers))
    fmt.Println("copy cnt:",copy(copyA,numbers))
    fmt.Println("copied data:",copyA)

    copyB :=make([]int,3)
    fmt.Println("copy cnt:",copy(copyB,numbers[2:5]))
    fmt.Println("copied data:",copyB)

    copyC :=make([]int,5)
    fmt.Println("copy cnt:",copy(copyC,numbers))
    fmt.Println("copied data:",copyC)


}