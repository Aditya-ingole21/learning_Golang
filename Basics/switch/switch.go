package main
import "fmt"
import "time"
func main() {
	// simple switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Another day")
	}

// switch with multiple cases
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
	}
// switch without expression
	typeOf := func (i interface{}){
		switch t := i.(type) {
		case int :
			fmt.Println("Type is int")
		case string:
			fmt.Println("Type is string")
		default:
			fmt.Println("Unknown type", t)

		}
	}
	typeOf(42)
	typeOf("Golang")
	typeOf(3.14)
}