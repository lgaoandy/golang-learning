package main 

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func userInput(prompt string) (answer string) {
	fmt.Print("\n", prompt)
	fmt.Scan(&answer)
	fmt.Print()
	return
}

func selectOption() (option int) {
	for {
		answer := userInput("Select an option: ")
		if isInteger(answer) {
			option, _ = strconv.Atoi(answer)
			return
		}
		fmt.Println("Invalid option.")
	}
}

func viewList() {
	f, err := os.Open("list.txt"); if err != nil {
		fmt.Println("Error opening file")
	}

	fmt.Println("Your list:")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("- %v\n", scanner.Text())
	}
}

func addToList() {
	// 
	// f, err := os.OpenFile("list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("Error opening file: %v\n", err)
    //     return
	// }

	// deter f.Close()

	// if _, err := f.WriteString()
}

func main() {
	fmt.Println("TO-DO LIST")
	fmt.Println("Here are your options:")
	fmt.Println("(1) - View List")
	fmt.Println("(2) - Add to List")
	fmt.Println("(3) - Exit")

	// User prompt menu
	action := selectOption()

	// Read from list
	switch action {
	case 1:
		viewList()
	case 2:
		addToList()
	}

}