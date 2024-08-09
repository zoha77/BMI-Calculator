package main

import (
	"bufio"
	"fmt"
	"myFirstApp/info"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Output information
	fmt.Println(info.BmiTitle)
	fmt.Println(info.Seperator)

	// Prompt for user input (weight + height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	weightInput = strings.TrimSpace(weightInput) // Use TrimSpace to remove whitespace
	heightInput = strings.TrimSpace(heightInput)

	// Parse the inputs to float64
	weight, err := strconv.ParseFloat(weightInput, 64)
	if err != nil {
		fmt.Println("Error parsing weight:", err)
		return
	}

	height, err := strconv.ParseFloat(heightInput, 64)
	if err != nil {
		fmt.Println("Error parsing height:", err)
		return
	}

	fmt.Println(weight)
	fmt.Println(height)

	// Calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your bmi: %.2f\n", bmi)
}
