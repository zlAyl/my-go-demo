package main

import (
	"sort"
	"strconv"
)

func SingleNumber(nums []int) []int {
	numberSet := make(map[int]int)
	for _, num := range nums {
		numberSet[num] += 1
	}
	var result []int
	for k, v := range numberSet {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

// IsPalindrome 判断数字是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	iToStr := strconv.Itoa(x)
	var newX string
	for k, _ := range iToStr {
		newX += string(iToStr[len(iToStr)-k-1])
	}

	if newX == iToStr {
		return true
	}
	return false
}

// IsValid 判断是否是有效的括号
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	strMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	for _, v := range s {
		right, isRight := strMap[v]
		if isRight == true { //表明v在key值里面 是右括号
			//检查栈是否为空或栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != right {
				return false
			}
			// 弹出匹配的左括号
			stack = stack[:len(stack)-1]
		} else { //表明v不在key值里面 是左括号 左括号压入栈里
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

// LongestCommonPrefix 最长公共前缀
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {

		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}

		prefix = prefix[:j]

		if prefix == "" {
			break
		}

	}

	return prefix
}

// PlusOne 加一
func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果加1后小于10，没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 否则，当前位设为0，进位到前一位
		digits[i] = 0
	}
	// 如果所有位都进位了，需要在数组前面添加1
	return append([]int{1}, digits...)

	//var str string
	//for _, v := range digits {
	//	str += strconv.Itoa(v)
	//}
	//
	//newInt, err := strconv.Atoi(str)
	//if err != nil {
	//	return nil
	//}
	//newInt += 1
	//
	//newStr := strconv.Itoa(newInt)
	//
	//digits = make([]int, len(newStr))
	//for i, v := range newStr {
	//	digit, _ := strconv.Atoi(string(v))
	//	digits[i] = digit
	//}
	//fmt.Println(digits)
	//return digits
}

func RemoveDuplicates(nums []int) (int, []int) {
	i := 0
	for j := 1; j < len(nums); j++ {
		// 当发现不重复的元素时
		if nums[j] != nums[i] {
			// 将不重复的元素移动到正确位置
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1, nums[:i+1]
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按区间的起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 如果当前区间与最后一个区间重叠，则合并
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 否则，将当前区间添加到结果中
			merged = append(merged, current)
		}
	}

	return merged
}

func TwoSum(nums []int, target int) []int {
	var sums []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 检查两数之和是否等于目标值
			if nums[i]+nums[j] == target {
				sums = append(sums, nums[i], nums[j])
			}
		}
	}
	return sums
}
