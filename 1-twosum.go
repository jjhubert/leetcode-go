package main

import "fmt"

/*
思路：通过两层循环前后前后相加去匹配是否等于目标数字
*/
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

/*
思路：
	1、先把切片转换为以值为key，下标为value的map
	2、通过把目标数字与迭代的数相减再判断map中是否存在该key
*/
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
