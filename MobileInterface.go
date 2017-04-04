package main

import "fmt"

type Phone struct {
	modelName string
	memory string
}

func (p Phone)Call()  {
	fmt.Println("Calling...!!")
}
func(p Phone)Sms(){
	fmt.Println("Sedding sms..!!")
}
type AndroidPhone struct {
	Phone
	AndroidVersion string
}

func main()  {

	pn:=Phone{modelName:"Nokia110",memory:"4Gb"}
	pn.Call()
	pn.Sms()
	and:=AndroidPhone{AndroidVersion:"4.5"}
	and=Phone{"samsung","4gb"}
	and.Sms()
	and.Call()

}