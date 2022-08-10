package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay, you can play!")
	} else {
		fmt.Println("I'm sorry. You're not quite old enough to play this game.")
		return
	}
	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have?")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You got %v out of %v questions correct. That's %v%%!", score, num_questions, percent)
}
