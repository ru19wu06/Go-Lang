package main

import(
    "fmt"
    "math"
)

type Point struct{
    x float64
    y float64
}

type Circle struct{
    r float64
    p Point
}

func main(){
	c := Circle{r:5.0,p:Point{x:3.0,y:4.0}}
	
	fmt.Println(c.p.x)
	fmt.Println(c.p.y)

	circumstance := 2*math.Pi*c.r

	fmt.Println(circumstance)

}

