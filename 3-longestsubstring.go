package main

import (
	"fmt"
)

/*
一、准备计数器与最大值int变量、保存字符出现的map
二、第一层循环迭代字符串
	1、计算器自增，因为字符出现
三、第二层循环下一个字符开始的字符串
	1、判断前一个字符跟当前字符是否相等，相等跳出循环
	2、判断字符字典即判断子字符串是否出现重复字符，当出现相同字符跳出循环，
	3、不相同就加计数器加1。添加以字符为key，下标为value的map
四、判断计数器与最大值哪个大，计数器大保存为最大值
*/
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	m := make(map[int]int)
	count := 0
	max := 0
	for i := 0; i < len(s) - 1; i++ {
		count++
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				break
			}
			if _, ok := m[int(s[j])];ok {
				break
			}
			m[int(s[j])] = j
			count++
		}
		if count > max{
			max = count
		}
		count = 0
		m = make(map[int]int)
	}
	if max == 0 {
		max = 1
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aaaaaaaa"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
}