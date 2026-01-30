package basics

import "fmt"

func main() {
	//defer example
	deferExample()

	//ellipses Examplefmt
	fmt.Println("Sum:", add(1, 2, 3, 4, 5))
	//passing slice to variadic function
	fmt.Println("Sum:", add([]int{90, 90, 9, 0, 89}...))

	//panic example with recover
	process()

	//panic example

	preprocess()

}
func deferExample() {
	defer fmt.Println("This will be printed last.")
	defer fmt.Println("heelo")
	fmt.Println("This will be printed first.")
	defer fmt.Println("This will be printed second.")
}
func add(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}
func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}

	}()
	fmt.Println("Processing...")
	panic("panic is called")
}
func preprocess() {
	defer fmt.Println("called")
	fmt.Println("Processing...")
	panic("panic is called")

}
