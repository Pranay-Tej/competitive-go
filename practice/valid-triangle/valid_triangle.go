package validtriangle

import "fmt"

func is_valid_triangle(sides []int) string {
	var total_length = len(sides)
	for target_side_index := range total_length {
		sum := 0
		for j := 0; j < total_length-1; j++ {
			idx := (target_side_index + j + 1) % total_length
			sum = sum + sides[idx]
		}
		if sum <= sides[target_side_index] {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	// 	var num_sides int;
	// 	fmt.Scan(&n)
	num_sides := 3
	var sides []int
	for range num_sides {
		var side int
		fmt.Scan(&side)
		sides = append(sides, side)
	}
	fmt.Println(is_valid_triangle(sides))
}
