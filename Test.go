package main
import  "TestProject"
import "fmt"


func main()  {
	var Result int =TestProject.AddNumber(10,20);
	fmt.Println("Addition of 10 and 20",Result)
	var ResultSub int =TestProject.SubNumber(100,20);
	fmt.Println("Subtraction of 100 and 20",ResultSub)
	var ResultMult int=TestProject.MultdNumber(10,5)
	fmt.Println("Multipication of 10 and 5",ResultMult)
}
