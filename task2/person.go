package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeId int
	Persons    Person
}

func (e *Employee) PrintInfo() {
	fmt.Println("员工ID为:", e.EmployeeId)
	fmt.Println("员工姓名为:", e.Persons.Name)
	fmt.Println("员工年龄为:", e.Persons.Age)
}

func main() {
	emp := Employee{
		EmployeeId: 1,
		Persons:    Person{Name: "张三", Age: 20},
	}
	emp.PrintInfo()
}
