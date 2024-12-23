package main // for an executable package, main the the major entry to run the program.
/*
In Go, the main package has a special significance.
It's used to define the entry point of a Go application.
Here's why you can (and sometimes must) use main as the package name instead of the actual logical package name:

It serves as the starting point for the program. When you run go run or build your program,
Go looks for the main package and specifically the func main() within it.

If your code is part of a library, the package name can reflect its purpose (e.g., utils, database).

However, for an executable program, you must have a main package because only this package is recognized by the Go toolchain as an entry point.

*/
import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
) // fmt packge contains functions for formatting text, including printing to the console.
func main(){
	fmt.Println("Hello, world!")

	fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	message,err := greetings.Hello("Gladys!")
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(message)
}
