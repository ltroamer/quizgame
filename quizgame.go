package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	askName()
	askAge()
}

func askName() {
	fmt.Print("Hi :) Please enter your name here: ")
	var name string
	name = runReadString()
	fmt.Println("Welcome to the game, " + name)
}

func askAge() {
	fmt.Print("How old are you: ")
	//	var age string
	age := runReadString()
	//	age = strconv.Atoi(input)
	fmt.Println("That's great, I always want to be " + age)
}

func runReadString() string {
	const inputdelimiter = '\n'
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
	}
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	return input
}
