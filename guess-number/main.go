package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const min int = 1
const max int = 100

var (
	colorRed   string = "\033[31m"
	colorGreen string = "\033[32m"
	colorBlue  string = "\033[34m"
	colorReset string = "\033[0m"
)

func main() {
	printHeading()
	for {
		play()
	}
}

func play() {
	secretNumber := generateSecretNumber()

	fmt.Println("> Greetings, stranger! I've just picked a number between 1 and 100. Try to guess it!")
	fmt.Print("> Please input your guess: ")

	attempts := 0
	for {
		attempts++

		fmt.Print(colorGreen)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		fmt.Print(colorReset)
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again.", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value. By the way, you just failed one of your attempts.")
			continue
		}

		switch {
		case guess > secretNumber:
			fmt.Print("> It is bigger than the secret number. Try again: ")
		case guess < secretNumber:
			fmt.Print("> It's smaller than the secret number. Try again: ")
		default:
			fmt.Printf("> Correct, stranger! You guessed right after "+colorRed+"%v"+colorReset+" attempts.\n\nLet's play one more time!\n\n", attempts)
			return
		}
	}
}

func generateSecretNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func printHeading() {
	fmt.Print(`
  ____                     _   _                 _                       _ 
 / ___|_   _  ___  ___ ___| \ | |_   _ _ __ ___ | |__   ___ ____  __   _/ |
| |  _| | | |/ _ \/ __/ __|  \| | | | | '_ ' _ \|  _ \ / _ \  __| \ \ / / |
| |_| | |_| |  __/\__ \__ \ |\  | |_| | | | | | | |_) |  __/ |     \ V /| |
 \____|\____|\___||___/___/_| \_|\__,_|_| |_| |_|____/ \___|_|      \_/ |_|
`)

	fmt.Println(colorBlue + `
               GuessNumber v1.0 by Akim Gubanov (c) 2020                    
` + colorReset)

	return
}

func init() {
	if runtime.GOOS == "windows" {
		colorReset = ""
		colorRed = ""
		colorGreen = ""
		colorBlue = ""
	}
}
