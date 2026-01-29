package basics

import "fmt"

type Employee struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	//PascalCase
	//E.g NewCustomer,GetItem,SetValue
	// Structs,Interfaces,enums

	//snake_case
	//E.g first_name,get_item,set_value
	//filenames,database fields

	//UPPERCASE
	//use case for Constants

	//mixedCase
	//htmlDocument,javaScript
	const MAXRETRIES = 5
	var employeeID = 1234
	fmt.Println("EmployeeID: ", employeeID)

}
