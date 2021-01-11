package main

//結構範例

import (
    "fmt"
)

type Person struct{
	Name string
}

type Saiyan struct{
	*Person
	Power int
}

func (p *Person) Inteoduce(){
	fmt.Printf("Hi , I'm %s\n",p.Name)
}

func main(){
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power: 9001,
	}

	goku.Inteoduce()//使用goku直接讀取Inteoduce()
}