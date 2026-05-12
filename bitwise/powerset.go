package bitwise

import (
	"fmt"
	"math"
)

func PowerSet(inputArr []int) {
	length := len(inputArr)

	for i := 0; i < int(math.Pow(2, float64(length))); i++ {
		answer := make([]int, 0)

		for j := 0; j < length; j++ {
			if i&(1<<j) != 0 {
				answer = append(answer, inputArr[j])
			}
		}

		fmt.Print("[")
		for j := 0; j < len(answer); j++ {
			fmt.Printf("%d", answer[j])
			if j != len(answer)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Print("]")
		fmt.Println()
	}
	fmt.Println()
}

// func main() {
// 	var t int
// 	fmt.Scan(&t)
// 	for t > 0 {
// 		var n int
// 		fmt.Scan(&n)
// 		arr := make([]int, n)
// 		for i := 0; i < n; i++ {
// 			fmt.Scan(&arr[i])
// 		}
// 		PowerSet(arr)

// 		t--
// 	}
// }
