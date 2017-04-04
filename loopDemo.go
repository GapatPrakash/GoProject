package main

import (
	"fmt"


)

func main()  {

	numbers:=[5]int{5,20,21,23,25}


	for index,val :=range numbers{
		fmt.Println(index,val)
	}

	 mySlice:=numbers[:]
	fmt.Println(mySlice)

	 numSlice:=make([]int,5,10)
	numSlice=append(numSlice,15,20)
	fmt.Println(numSlice)

}
