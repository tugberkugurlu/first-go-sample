package main

import "fmt"

func main() {
	message := "Hello Go World!"
	
	// greeting is a pointer and we are assigning the memory address of the message variable
	var greeting *string = &message
	
	// will print something like: Hello Go World! 0xc0820001c0
	fmt.Println(message, greeting)
	
	// if we do this, this will do the oposite of the dereferencing operator (&) and gets the value
	fmt.Println(message, *greeting)
	
	*greeting = "hi"
	
	// should print out "hi hi"
	fmt.Println(message, *greeting)
}