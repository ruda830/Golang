package main

import (
	"fmt"

	"example.com/greetings"
)
/*Install two package to call the functions.
fmt and
example.com/greetings mod,That is, the package included in the example.com/greetings module (in this case, the greetings package)
*/

func main() {
	// Get a greeting message and print it.
	//Access the Hello function in the greetings package
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
