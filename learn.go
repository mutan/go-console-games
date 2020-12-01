package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %T %v\n", input, input)
	fmt.Println("\033[2J")
}
