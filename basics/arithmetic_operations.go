package basics

func main() {
	var a, b int = 10, 20
	var result int
	result = a + b
	println("Addition: ", result)

	result = a - b
	println("Subtraction: ", result)

	result = a * b
	println("Multiplication: ", result)

	result = b / a
	println("Division: ", result)

	result = b % a
	println("Modulus: ", result)

	const pi float32 = 22.0 / 7
	println("Value of Pi: ", pi)
}
