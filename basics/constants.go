package basics

import "fmt"

const PI = 3.14
const GRAVITY = 9.8

func main() {
	const days int = 7
	const (
		MONDAY        = 1
		TUESDAY       = 2
		WEDNESDAY int = 3
	)

	fmt.Println("Value of PI:", PI)
	fmt.Println("Value of GRAVITY:", GRAVITY)
	fmt.Println("Days in a week:", days)
	fmt.Println("Monday is day number:", MONDAY)
	fmt.Println("Tuesday is day number:", TUESDAY)
	fmt.Println("Wednesday is day number:", WEDNESDAY)

}
