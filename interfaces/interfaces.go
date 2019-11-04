package main

import "fmt"


type geometry interface {
	area() float64 //method signature: methodname(parameter types) return type i.e takes an int and returns a bool
	perim() float64
}

type rect struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.height*r.width;
}

func (c circle) area() float64{
	return 3.14*c.radius*2;
}

func (r rect) perim() float64{
	return 2*r.height+ 2*r.width;
}

func (c circle) perim() float64{
	return c.radius*2;
}

func measure(g geometry){
	fmt.Println("Geometry: ",g)
	fmt.Println("Area:" , g.area())
	fmt.Println("Perimeter", g.perim())
}

func main(){
	fmt.Println("This is interfaces go program")

	r:= rect{width:  10,height: 20}
	c:=circle{radius:5}

	measure(r)
	measure(c)

}