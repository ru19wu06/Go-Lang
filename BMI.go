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
	Marry := Figure{Weight:58,Hight:160}
	His_BMI := Aaron.Weight/(((Aaron.Hight/100)*(Aaron.Hight/100)))
	Her_BMI := Marry.Weight/(((Marry.Hight/100)*(Marry.Hight/100)))

	L := BMI{bmi:His_BMI,F:Aaron}
	M := BMI{bmi:Her_BMI,F:Marry}
	fmt.Println("Aaron BMI :")
	fmt.Println(L.bmi)

	if L.bmi>30 {
		fmt.Println("太胖")
	}else{
		fmt.Println("太瘦")
	}


	fmt.Println("Marry BMI :")
	fmt.Println(M.bmi)

	if M.bmi>30 {
		fmt.Println("太胖")
	}else if M.bmi <20{
		fmt.Println("太瘦")
	}else{
		fmt.Println("均勻")
	}

}