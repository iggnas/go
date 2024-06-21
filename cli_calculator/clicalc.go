package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanNumber() float64 {
	fmt.Print("\nchoose number: ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple CLI GoLang calculator")
	fmt.Print("\nstarting number")
	f := scanNumber()

	for {
		fmt.Print("\nprompt: ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		switch scanner.Text() {
		case "h", "help":
			fmt.Println("\n+ for addition")
			fmt.Println("\n- for subtraction")
			fmt.Println("\n* for multiplication")
			fmt.Println("\n/ for division")
			fmt.Println("\nv for current value")
			fmt.Println("\nq to quit")
		case "+":
			n := scanNumber()
			f = f + n
			fmt.Println("Result:", f)
		case "-":
			n := scanNumber()
			f = f - n
			fmt.Println("Result:", f)
		case "*":
			n := scanNumber()
			f = f * n
			fmt.Println("Result:", f)
		case "/":
			n := scanNumber()
			if n != 0{
				f = f / n
				fmt.Println("Result:", f)
			}else{
				fmt.Println("Canot divide by 0")
			}
		case "v":
			fmt.Println(f)
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Type 'h' for help.")
		}
	}

}
