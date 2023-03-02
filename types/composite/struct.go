package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	employeeOne := employee{"John", 21, 1000}
	fmt.Println(employeeOne)

	employeeTwo := employee{
		name:   "Sam",
		age:    22,
		salary: 1100,
	}
	fmt.Println(employeeTwo)

	employeeThree := employee{name: "Tina", age: 24}
	fmt.Println(employeeThree)

  append()
}
