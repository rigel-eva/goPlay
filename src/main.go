package main

import (
	"fmt"

	"./hello" //Ok this is how we reference a package we wrote ourself in our code!
	"./loopsAndControlFlow"
	"./typesAndClasses"
)

/* Ok sooo ...
 *  Capitals are important for go, they denote whether or not a functions/variable should be "exported" (made public) or should be "unexported" (made private)
 */
func main() {
	fmt.Print("\nPackage hello: \n\n")
	hello.Main()
	fmt.Print("\nPackage loopsAndControlFlow: \n\n")
	loopsAndControlFlow.Main()
	fmt.Print("\nPackage typesAndClasses: \n\n")
	typesAndClasses.Main()
}
