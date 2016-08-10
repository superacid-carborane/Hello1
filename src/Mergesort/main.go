package main

import "fmt"

func mergesort(m []int) []int {
	fmt.Println(m)
	if len(m) < 2 {
		return m
	}
	middle := len(m) / 2
	left := m[:middle]
	right := m[middle:]

	left = mergesort(left)
	right = mergesort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; i < len(result); i++ {
		for i < len(left) && i < len(right) {
			if i < len(left) && i < len(right) {
				if left[i] < right[i] {
					result = append(result, left[i])
					//fmt.Println(result)
					i++
					fmt.Println(i)
				} else {
					result = append(result, right[i])
					//fmt.Println(result)
					i++
					fmt.Println(i)
				}
			} else if left[i] < right[i] {
				result = append(result, left[i])
				//fmt.Println(result)
				i++
				fmt.Println(i)
			} else if right[i] < left[i] {
				result = append(result, right[i])
				//fmt.Println(result)
				fmt.Println(i)
				i++
			} else {
				break
			}
		}
	}
	return result
}

func main() {
	test_slice := []int{2, 1, 6, 8, 3, 9, 11, 22}
	fmt.Println(mergesort(test_slice))
}
