package main

import "fmt"
func  main()  {

    syaHello:= func(name string) string {
        fmt.Print("Hello ")
        return name
    }

    f11()
    defer f3()
    f12()
  fmt.Print(syaHello("Prakash\n"))
}
func f11()  {
    fmt.Printf("\nIn F11")
}
func  f3()  {
    fmt.Print("\nIn f3")

}
func f12()  {
    fmt.Print("\nIn f12")

}