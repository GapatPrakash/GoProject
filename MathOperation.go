package main

import (
	"fmt"

)

func main()  {

	AdditionResult:=addNumber(20,80)
	fmt.Println(AdditionResult)

	SubResult:=subNumber(82,20)
	fmt.Println(SubResult)

	lengthOfString:=LengthOfString("Prakash")
	fmt.Println(lengthOfString)

}
func addNumber(num1, num2 int)(result int)  {
	fmt.Println(result)
	result= num1+num2;
	return

}
func  subNumber(num1, num2 int)(int)  {
	return num1-num2;
}
func LengthOfString(name string)(int){

	len:=len(name)
	return len
}