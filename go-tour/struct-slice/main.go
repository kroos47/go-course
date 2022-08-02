package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func point(){
	i:=1
	p := &i
	fmt.Println(*p)
	*p = 2
	fmt.Println(i)
}

type Rectangle struct{
	X, Y float64
}

func slices(){
	s:= []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(s)
	s = s[:0]
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
	s = s[1:4]
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
	s = s[:8]
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
	s = s[2:]
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)

}

func Pic(dx, dy int) [][]uint8 {
	arr:= make([][]uint8,dy)
	for i := range arr{
		arr[i] = make([]uint8, dx)
		for j := range arr[i]{
			arr[i][j] = uint8(i+j)
		}
	}
	return arr
}
func main(){
	point()	

	x1 := Rectangle{1,2}
	x2 := Rectangle{}
	p := &Rectangle{4,5}
	fmt.Println(x1,x2,p,p.X)

	// var a [2]string
	arr :=[6]int{1,2,3,4,5,6}
	fmt.Println(arr)
	sli := arr[1:5]
	fmt.Println(sli)
	sli[2] = 10 
	fmt.Println(arr)

	slices()

	var app []int
	fmt.Printf("len = %d cap = %d %v \n", len(app), cap(app), app)
	app = append(app, 1)
	fmt.Printf("len = %d cap = %d %v \n", len(app), cap(app), app)
	app = append(app, 2,3,4,5,6,7,8,9,10)
	fmt.Printf("len = %d cap = %d %v \n", len(app), cap(app), app)

	for _,val := range app {
		fmt.Println(val)
	}

	pic.Show(Pic)

}