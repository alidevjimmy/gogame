package move

import "fmt"

// GoForward move unit to specific position
func GoForward(position int) {

	for i := 0; i <= position; i++ {
		fmt.Println("cell :", i)
	}

}
