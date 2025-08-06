package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	//Set properties of the predefiened Logger, including the log entry prefix
	//and a flag to disable printing the time, source file and a line number
	log.SetPrefix("greetings: ")
	log.SetFlags(2)
	names := []string{"Goder", "Gladys", "Goodra"}
	message, err := greetings.Hellos(names)
	if err != nil {
		//If an error was returned, print it to the console and exit the program
		log.Fatal(err)
	}

	fmt.Println(message)
}
