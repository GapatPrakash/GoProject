package main

import "fmt"
import "errors"

func main(){
returnvale,error1 :=myFunction(0);
}
func myFunction(i int)(int error){
	if i<0{
		return -1,errors.New("Number should be greater than 0")
	}
	return i,nil
}


