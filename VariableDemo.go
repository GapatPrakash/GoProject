package main

import "fmt"

const CarName  = "Hey"
var (
	carModel="odi"
	carManufectureYear="2010"
	carPrice=2000000
)
	func main(){
		fmt.Println(CarName)
		fmt.Print("Car Deatils are as bellow\n")
		fmt.Println(carModel)
		fmt.Println(carManufectureYear)
		fmt.Println(carPrice)

		var name,id = DisplayMessage("Prakash",4604)
		fmt.Println(name,id)

	}

	func DisplayMessage(FirstName string, EmployeeCode int)(string, int){

		return FirstName,EmployeeCode;
	}

