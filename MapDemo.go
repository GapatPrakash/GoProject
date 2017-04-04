package main

import (
	"fmt"


)

func main(){

	CityData:=map[string]string{
		"OS":"Osmanabad",
		"L":"Latur",
		"B":"Beed",
	}
	Display(CityData)


	result:=GetDataByKey(CityData,"OS1")
	fmt.Println(result)

}
func Display(data map[string]string){

	for _,val:=range data{
		fmt.Println(val)
	}
}
func GetDataByKey(data1 map[string]string,key string )(string){
	fmt.Println("value for the ",key+" \t is")
	for key1,val:=range data1{
		if key1==key {
			return val
		}


	}
	return "Not Found"
}