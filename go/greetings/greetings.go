/*Return the greeting to the person who greeted you.*/
package greetings

import "fmt"

//Hello returns a greeting for named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

/* := operator is a shortcut for declaring and initializing variables. It is the same as the following.
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
*/

/*Sprintf is a function that returns a greeting statement, replacing the name parameter value with the %v format.
*/