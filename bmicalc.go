package main

import (
	"fmt"
)

func main() {
	fmt.Println("\t\t\tWELCOME TO BMI CALCULATOR\t\t\t")
	fmt.Println("\t\t\t=========================\t\t\t")

	var weight, height float64

	fmt.Println("Enter your weight (kgs): ")
	fmt.Scanln(&weight)
	fmt.Println("Enter your height (m): ")
	fmt.Scanln(&height)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is: %f\n", bmi)

	//Underweight: bmi < 18.5
	//Normal weight: bmi 18.6 -> 24.5
	//overweight: bmi 34.6 -> 29.9
	//Obesity: bmi >= 30

	if bmi <= 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.6 && bmi <= 24.5 {
		fmt.Println("Category: Normal Weight")
	} else if bmi >= 34.6 && bmi <= 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}
