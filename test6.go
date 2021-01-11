package main
 
import (
    "fmt"
)
 
type Gender int
 
const (
    Male Gender = iota
	Female
	Tas
)
 
type Person struct {
    name   string
    gender Gender
    age    uint
}


func main() {
	me := Person{gender: Male, age: 30, name: "Michael Chen"}
	Thet := Person{gender: Tas, age:55,name: "A group of people"}
	he := Person{gender: Female, age:66, name : "Dasia tona"}
 
    fmt.Println(me.name)
    fmt.Println(me.age)
 
    if me.gender == Male {
        fmt.Println("Male")
    } else {
        fmt.Println("Female")
	}
	fmt.Println("--------------------------")
	fmt.Println(he.name)
	fmt.Println(he.age)

	if he.gender == Male {
        fmt.Println("Male")
    } else if he.gender == Female {
        fmt.Println("Female")
	}
	fmt.Println("--------------------------")
	fmt.Println(Thet.name)
	fmt.Println(Thet.age)

	if Thet.gender == Male {
        fmt.Println("Male")
    } else if Thet.gender == Female {
        fmt.Println("Female")
	}else{
		fmt.Println("else")
	}


}