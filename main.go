package main

import "fmt"

func main() {

	nums := []int{1, 3}

	target := 1
	result := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target && (len(result) == 0 || len(result) == 1) {
			result = append(result, i)
		}
		if nums[i] == target && len(result) == 2 {
			result[1] = i
		}

	}

	if len(result) == 1 {
		result = append(result, result[0])
	}
	if len(result) == 0 {
		result = append(result, -1, -1)
	}
	fmt.Println(result)
}
