package main

import (
	"fmt"
)


var z = 100.0 // the seed
var delta = 0.0001 // the epsilon

func Sqrt(x float64) float64 {	
	z -=  (z*z - x) / (2 * z)
	if z*z-x < delta {
		return z
	}
	fmt.Printf("value of z = %f \n", z)
	defer fmt.Printf("value of z = %f \n", z)
	return Sqrt(x)
}
func main(){

	for i:=0;i<=10;i++{
		fmt.Println(i)
	}
	sum:=0
	for sum < 10 {
		fmt.Println(sum)
		sum += 1
		if e:=2; sum % e == 0{
			sum += 1
		}
	}
	z:=(Sqrt(2))
	fmt.Println(z)

	swi := 5 
	switch swi {
	case 1: fmt.Println("one")
	case 2: fmt.Println("two")
	case 3: fmt.Println("three")
	case 4: fmt.Println("four")
	case 5: fmt.Println("five")
	default: fmt.Println("unknown")

	}
}   