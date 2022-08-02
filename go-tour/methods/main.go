package main 
import "fmt"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main(){
	c := circle{radius: 5}
	fmt.Println(c.area())
}