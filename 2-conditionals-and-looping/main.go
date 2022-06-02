package main

import (
	"fmt"
	"os"
)

func main() {

	var option int

	fmt.Println("Enter option:")
	fmt.Scanf("%d", &option)

	//diferente de C ele espera um boolean no if, não considerando números
	if option == 1 {
		fmt.Println("Option 1")
	} else if option == 2 {
		fmt.Println("Option 2")
	} else if option == 3 {
		fmt.Println("Option 3")
	}

	//não precisa colocar o break no switch em go
	switch option {
	case 1:
		fmt.Println("Option 1")
	case 2:
		fmt.Println("option 2")
	case 3:
		fmt.Println("Option 3")
	default:
		fmt.Println("Undefined")
	}

	for {
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Println("Option 1")
		case 2:
			fmt.Println("option 2")
		case 3:
			fmt.Println("Option 3")
		default:
			fmt.Println("Undefined")
			os.Exit(1)
		}
	}
}
