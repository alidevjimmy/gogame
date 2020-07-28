package move

import "fmt"
// GoLeft move unit to specific position
func GoLeft(position int) {
	for i := position; i >= 0; i-- {
		fmt.Println("cell :", i)
	}
}
