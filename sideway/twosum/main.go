package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//for i := 0; i < len(nums); i++ {
	//	for j := 1; j < len(nums); j++ {
	//		if j < len(nums) {
	//			if ((nums[i] + nums[j]) == target) && (i != j) {
	//				return []int{i, j}
	//			}
	//		}
	//	}
	//}
	//return nil
	var output []int
	//convert slice to map
	numsMap := make(map[int]int)
	for key, value := range nums {
		numsMap[value] = key
	}
	//find in map
	for key, value := range nums {
		result := target - value
		nextKey, exist := numsMap[result]
		if exist && nextKey != key {
			output = append(output, key, nextKey)
			break
		}
	}

	return output
}

func main() {
	nums := []int{2, 5, 5, 11}
	//target := 10

	numsMap := make(map[int]int)
	for key, value := range nums {
		numsMap[value] = key
	}
	index, exist := numsMap[11]
	fmt.Println(index, exist)
	//fmt.Println(twoSum(nums, target))
}
