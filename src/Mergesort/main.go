package main

import "fmt"

func mergesort(m []int) []int {
	fmt.Println(m)
	if (len(m)) < 2 {
		return m
	}
	middle := (len(m)) / 2
	left := m[:middle]
	right := m[middle:]

	left = mergesort(left)
	right = mergesort(right)
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	left_tmp, right_tmp := 0, 0
	for left_tmp <= len(left)-1 && right_tmp <= len(right)-1 {
		if left[left_tmp] <= right[right_tmp] {
			result = append(result, left[left_tmp])
			fmt.Println(result)
			left_tmp += 1
		} else {
			result = append(result, right[right_tmp])
			fmt.Println(result)
			right_tmp += 1
		}
	}
	return result
}

func main() {
	test_slice := []int{1, 2, 6, 8, 3, 9, 11, 22, 4, 36}
	fmt.Println(test_slice)
	fmt.Println(mergesort(test_slice))
}
