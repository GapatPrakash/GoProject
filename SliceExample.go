package main

import (
	"fmt"

)

func main()  {

	Numbers :=make([]int,5,10)
	Numbers[0]=10;
	Numbers[1]=20;

	for val,val2 :=range Numbers{
		if val2==20 {
			continue
		}else{
			fmt.Println(val,val2)
		}
	}

	fmt.Println("New===>")
	data:=make(map[int]string)
	m:=map[string]string{
		"Prakash":"Gapat",
		"Piyush":"Mulik",

	}
	fmt.Println(m)


	fmt.Println(Numbers)
	MyFunction(Numbers,1)
	fmt.Println(Numbers)

}
func MyFunction(x []int,val int){
	x[val]=val*2;
}
