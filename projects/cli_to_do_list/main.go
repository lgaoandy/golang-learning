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
		fmt.Println("\nOptions Menu:")
		fmt.Print("(1) View List, ")
		fmt.Print("(2) Add to List, ")
		fmt.Print("(3) Exit")
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

	fmt.Println("\nYour list:")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("- %v\n", scanner.Text())
	}
}

func addToList() {
	newTask := userInput("New task: ")
	
	f, err := os.OpenFile("list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
        return
	}

	defer f.Close()

    if _, err := f.WriteString(newTask + "\n"); err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
        return
    }
}

func main() {
	fmt.Println("TO-DO LIST")

	// User prompt menu
	var action int

	for {
		action = selectOption()

		// Read from list
		switch action {
		case 1:
			viewList()
		case 2:
			addToList()
		case 3:
			return
		}
	}
}