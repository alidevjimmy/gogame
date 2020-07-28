package move

import "fmt"

// GoRight move unit to specific position
func GoRight(position int) {

	for i := 0; i <= position; i++ {
		fmt.Println("cell :", i)
	}

}
