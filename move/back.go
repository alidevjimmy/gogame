package move

import "fmt"
// GoBack move unit to specific position
func GoBack(position int) {
	for i := position; i >= 0; i-- {
		fmt.Println("cell :", i)
	}
}
