package main 

import (
	"fmt"
	"math"
)

func addition(x,y int) int {
	return x + y
}
func swapping(x, y, z string) (string, string, string) {
	return z,y,x
}
func naked_return(sum int) (x,y int) {
	x = sum * 4+1
	y = sum * 5+2
	return
}
var z, python, java = true, false, "yeh!"

func main(){
	fmt.Println(math.Pi)

	fmt.Println(addition(1,2))

	a, b, c := swapping("I","am","testing")
	fmt.Println(a,b,c)
	fmt.Println(swapping("I","am","Kailas"))

	fmt.Println(naked_return(5))

	var i int
	fmt.Println(i, z, python, java)

	var v,k,l = 1,2,3
	fmt.Println(v,k,l,python,java)

	d := 5
	test, test2 := true , "test2" 
	fmt.Println(d,test,test2)

	fmt.Printf("Type: %T Value: %v\n", test, test)

	// var i1 int
	var i2 string
	// var i3 bool
	fmt.Printf(" %q \n", i2)

	var x,y int = 5,3
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)

	s:=1
	t:= float64(i)
	fmt.Printf("%T %T",s,t)
	

}