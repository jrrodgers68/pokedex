package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := NewConfig()
	for {
		printPrompt()

		scanner.Scan()

		tokens := strings.Fields(strings.ToLower(scanner.Text()))

		err := invokeCommand(tokens[0], tokens[1:], config)
		if err != nil {
			fmt.Println(err)
			//break
		}
	}
}
