package main

import(
    "fmt"
    
)

type Figure struct{
	Weight float64
	Hight float64
}

type BMI struct{
	bmi float64
	F Figure
}

func main(){
	Aaron := Figure{Weight:98,Hight:177}
	His_BMI := Aaron.Weight/(((Aaron.Hight/100)*(Aaron.Hight/100)))

	L := BMI{bmi:His_BMI,F:Figure{Weight:98,Hight:177}}

	fmt.Println("Aaron BMI :")
	fmt.Println(L.bmi)

	if L.bmi>30 {
		fmt.Println("太胖")
	}else{
		fmt.Println("太瘦")
	}

}