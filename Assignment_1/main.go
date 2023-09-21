package main

import (
	"assignment_1/Engineer"
	"fmt"
)

func main() {
	engineer := Engineer.NewEngineer()
	engineer.SetPosition("Boss")
	//fmt.Println(engineer.GetPosition())
	engineer.SetSalary(1000000)
	//fmt.Println(engineer.GetSalary())
	engineer.SetAddress("Baker street")
	//fmt.Println(engineer.GetAddress())
	fmt.Println(engineer)
}
