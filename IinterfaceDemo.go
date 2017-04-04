package main

import "fmt"

type myInterface interface {
	PrintArea()
}
type Squre struct {
	length int
}

func (s Squre) PrintArea(){
	fmt.Println("Area of squre of its  side  ",s.length," is ", s.length*s.length)
}
func main(){
	var s myInterface
	s=Squre{5}
	s.PrintArea()


}