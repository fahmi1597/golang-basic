package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// nums = [2,7,11,15], target = 9
	// return the indices, not the value of array
	// store it on result : {nums[value] : nums[index]}
	result := make(map[int]int)

	//iterate over array
	for index, value := range nums {
		// value = nums[index]
		// is target - value available in result?
		// example :
		// iterate 1, is 9 - 2 = 5, is 5 available in result?
		// iterate 2, is 9 - 7 = 2, is 2 available in result?

		if _, isExist := result[target-value]; isExist {
			// if yes, return the result[target-value] and current index
			// iterate 2, result[2] --> index 0 is available| 2 --> 0
			// then return result[2] --> index 0, and current index which is 1.
			return []int{result[target-value], index}
		}
		// if no, store current index in result[value], iterate to the next
		// iterate 1, store current index which is 0 in result[2] | 2 --> 0
		result[value] = index
	}
	return nil

}

func main() {
	arr := []int{2, 7, 11}
	fmt.Println(twoSum(arr, 9))

}
