package main

import "fmt"

func FindIndicesTwoNumbers(nums []int, target int) []int {
	numlen := len(nums)

	for i := 0; i < numlen; i++ {
		for j := i + 1; j < numlen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	fmt.Println(FindIndicesTwoNumbers([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
