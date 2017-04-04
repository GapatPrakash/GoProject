package main

import "fmt"

func main() {

	f1()

	f2(10, "Go")

	result := addTwoNumbers(10, 90)
	fmt.Println(result)
	fmt.Println(addTwoNumbers(10, 90))

	fullname, age := f4("Rob","Pike",60)
	fmt.Println(fullname, age)
	//fullname, _ := f4("Rob","Pike",60)
	//fmt.Println(fullname)
	//_, age:= f4("Rob","Pike",60)
	//fmt.Println(age)



	//Inline Functions
	//function to add two numbers
	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

	//function to add variable count of numbers
	addNumbers := func(numbers ...int) int {

		total := 0
		for i := 0; i < len(numbers); i++ {
			total += numbers[i]
		}
		return total

	}

	//function to return a greeting
	sayHello := func(name string) string {
		return "Hello " + name
	}

	fmt.Println(add(10, 20))
	fmt.Println(addNumbers(10, 20, 30, 40))

	fmt.Println(sayHello("Go"))
	fmt.Println(sayHello("Python"))

}

//No input parameters, no return values
func f1() {
	fmt.Println("I am f1")
}

//Multiple input parameters, no return values
func f2(num1 int, name string) {
	fmt.Println("I am f2")
	fmt.Println("The number is", num1)
	fmt.Println("The name is", name)
}

//Multiple input parameters, single return value
func addTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

//Multiple input parameters, Multiple return values
func f4(firstName string, lastName string, age int) (string, int) {
	return firstName + " " + lastName, age
}

//Named Parameters
func f5(firstName string, lastName string, age int) (fullname string, newage int) {
	fmt.Println("In f5")
	fmt.Println("Value of fullname:", fullname)
	fmt.Println("Value of newage:", newage)
	fullname = firstName + " " + lastName
	age += 10
	return
}

func addLotsOfNumbers(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

