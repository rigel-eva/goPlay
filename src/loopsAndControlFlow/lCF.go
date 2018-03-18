package loopsAndControlFlow

import "fmt"

func Main() {
	for i := 1; i < 20; i++ {
		fmt.Printf("Loop %v\n", i)
	}
	if true {
		fmt.Println("Oh look it's true!")
	} else {
		fmt.Println("Aww ... it's false ....")
	}
	//Oh hey it's a comment, neato toleto!
	whileLoops(5)
}
func whileLoops(initialValue int) { //ok that's a weird way to pass in values 🤷
	var x = initialValue
	var y = &x
	fmt.Printf("X Value: %v\n", x)
	fmt.Printf("X in go format: %#v\n", x)
	fmt.Printf("X's type: %T\n", x)
	fmt.Printf("X in base 16: %x\n", x)
	fmt.Printf("X in base 16 a different way: %X\n", x)
	fmt.Printf("Pointer to X: %p\n", &x)
	fmt.Printf("Y's Type: %T\n", y)
	fmt.Printf("Where does y point?: %p\n", y)
	//Ok and now time for some fun~
	//So Go has no while loops ... but their for loops are the while loops!
	for *y > 0 { //Ok we are watching where our pointer points
		fmt.Printf("X: %v Y: %v\n", x, *y)
		x-- //Which we are modifing where it's pointing here
	}
}
