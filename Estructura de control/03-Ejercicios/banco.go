package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter your age:")
	var age int
	fmt.Scanln(&age)
	fmt.Println("you are an employee (true/false):")
	var isEmployed bool
	fmt.Scanln(&isEmployed)
	var employmentOverOneYear bool
	fmt.Println("Have you been working in this company for more than a year? (true/false):")
	fmt.Scanln(&employmentOverOneYear)
	fmt.Println("What is your income:")
	var income int
	fmt.Scanln(&income)
	if age > 22 && isEmployed && employmentOverOneYear {
		if income > 100000 {
			fmt.Println("Congratulations, you can get an interest-free loan.")
		} else {
			fmt.Println("Congratulations, you can get a loan.")
		}
	} else {
		fmt.Println("Sorry, you don't qualify for a loan.")
	}
}
