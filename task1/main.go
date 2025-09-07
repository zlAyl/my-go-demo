package main

import (
	"fmt"
)

func main() {
	//go基础-----------task1
	//1.只出现一次的数字
	nums := []int{1, 1, 2, 3, 3, 6, 2, 5, 5, 6, 7}
	result := SingleNumber(nums)
	if result == nil {
		fmt.Println("没有只出现一次的数字")
	} else {
		fmt.Println("只出现一次的数字有:", result)
	}

	//2.回文数
	x := -1221
	result2 := IsPalindrome(x)
	if result2 == false {
		fmt.Printf("数字 x = %d 不是回文数\n", x)
	} else {
		fmt.Printf("数字 x = %d 是回文数\n", x)
	}

	//3.有效的括号
	str := "([])[]{"
	result3 := IsValid(str)
	if result3 == false {
		fmt.Printf("字符串 str = %s 不是有效的括号\n", str)
	} else {
		fmt.Printf("字符串 str = %s 是有效的括号\n", str)
	}

	//4.最长公共
	strs := []string{"flower", "flow"}
	result4 := LongestCommonPrefix(strs)
	fmt.Printf("strs 最长公共前缀为 %s \n", result4)

	//5.加一
	digits := []int{1, 2, 9}
	result5 := PlusOne(digits)
	if result5 == nil {
		fmt.Println("digits 加一后 处理错误 ")
	} else {
		fmt.Println("digits 加一后的值为 ", result5)
	}

	//6.删除有序数组中的重复项
	nums = []int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7}
	lens, result6 := RemoveDuplicates(nums)
	fmt.Println("删除有序数组中的重复项后数组的长度：", lens)
	fmt.Println("删除有序数组中的重复项后的数组为:", result6)

	//7.合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {3, 7}, {8, 9}}
	result7 := Merge(intervals)
	fmt.Println("合并后为:", result7)

	//8.两数之和
	nums = []int{2, 7, 11, 15}
	target := 17
	result8 := TwoSum(nums, target)
	fmt.Println("两数之和得到目标值得两个数为:", result8)
}
