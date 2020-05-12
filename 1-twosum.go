package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var resSli []int
	for i := 0; i < len(nums); i++ {
		for j:= i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				resSli = append(resSli, i)
				resSli = append(resSli, j)
			}
		}
	}
	return resSli
}

func twoSumPlus(nums []int, target int) []int {
	resMap := map[int]int{}
	for i, v := range nums {
		resMap[v] = i
	}
	for i, v := range nums {
		tmpNum := target - v
		if _, ok := resMap[tmpNum]; ok && i != resMap[tmpNum]{
			return []int{i, resMap[tmpNum]}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{11, 7, 2, 15}, 9))
	fmt.Println(twoSum([]int{2, 15, 11, 7}, 9))
	fmt.Println(twoSumPlus([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumPlus([]int{11, 7, 2, 15}, 9))
	fmt.Println(twoSumPlus([]int{2, 15, 11, 7}, 9))
	fmt.Println(twoSumPlus([]int{3,2,4}, 6))
}
