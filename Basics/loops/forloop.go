package main

import "fmt"

func main() {
// for loop with only condition
i := 1
for i<=3{
	fmt.Println(i)
	i++
}
// traditional for loop
for i=0; i<=5;i++{
	fmt.Println(i)
}
// for range loop
for i := range 4{
	fmt.Println(i)
}
}