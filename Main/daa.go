package Main

package aceo_077

import (
"fmt"
)

func main() {
	var boxes int
	fmt.Println("Amount of boxes:")
	fmt.Scanf("%s\n", &boxes)
	var array_of_boxes []int
	array_of_boxes = make([]int, boxes)

	for j := 1; j < boxes+1; j++ {
		for i := 0; i < boxes; i++ {
			if (i+1)%j == 0 {
				if array_of_boxes[i] == 0 {
					array_of_boxes[i] = 1
				} else {
					array_of_boxes[i] = 0
				}
			}
			fmt.Printf(string(array_of_boxes[i]))
		}
		fmt.Println(" ")
	}
}
