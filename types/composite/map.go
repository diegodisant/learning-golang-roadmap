package main

import "fmt"

func main() {
	var employeeSalary map[string]int

	fmt.Println(employeeSalary)

	var employeeSalary2 = make(map[string]int)

	fmt.Println(employeeSalary2)

	employeeSalary3 := map[string]int{
		"John": 1000,
		"Sam":  1000,
	}

	fmt.Println(employeeSalary3)
}
