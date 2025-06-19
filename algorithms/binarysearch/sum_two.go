package binarysearch

import "fmt"

// given an array of sorted integers and a sum
// return indices of two elements which will add up to given sum
// if not possible return -1

func rightBinSearch(numbers []int, index int, searchValue int) (int, bool) {
	start := index
	end := len(numbers) - 1
	for {
		if start > end {
			return -1, false
		}
		mid := (start + end) / 2
		val := numbers[mid]
		if val == searchValue {
			return mid, true
		}
		if val < searchValue {
			start = mid + 1
		}

		if val > searchValue {
			end = mid - 1
		}
	}

}

func sumTwoImpl(numbers []int, sum int) (int, int, error) {
	for i, num := range numbers {
		val, ok := rightBinSearch(numbers, i, sum-num)
		if ok {
			return i, val, nil
		}
	}
	return 0, 0, fmt.Errorf("not found")

}

// func sumTwo() {
// 	var length int
// 	fmt.Scan(&length)
// 	var numbers []int
// 	for range length {
// 		var num int
// 		fmt.Scan(&num)
// 		numbers = append(numbers, num)
// 	}
// 	var sum int
// 	fmt.Scan(&sum)

// 	// fmt.Println(numbers, sum)
// 	left, right, err := sumTwoImpl(numbers, sum)
// 	if err != nil {
// 		fmt.Println(-1)
// 		return
// 	}
// 	fmt.Printf("%d %d\n", left, right)
// }

// func main() {
// 	sumTwo()
// }
