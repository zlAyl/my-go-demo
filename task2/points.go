package main

import "fmt"

func SumTen(num *int) {
	*num += 10
}

func SliceTimes(numsPtr *[]int) {
	nums := *numsPtr
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
}

func main() {
	num := 2
	SumTen(&num)
	fmt.Println(num)

	nums := []int{2, 3, 4, 5}
	SliceTimes(&nums)
	fmt.Println(nums)
}
