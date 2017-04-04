package main

import (
	"fmt"

)

type  Emoloyee struct {
	 FirstName string
	 LastName string
	 Age int
	Address Address
}

type Address struct {
	CityName string
	PinCode int
}

func main(){

	Emp:=Emoloyee{FirstName:"Prakash",LastName:"gapat",Age:23}
	Emp.Address=Address{CityName:"Indapur",PinCode:413525}
	fmt.Println(Emp)


}

