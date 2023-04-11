package main

import (
	"fmt"
)

func counter(i int) {
	if i == 0 {
		return
	}

	fmt.Print(i, " ")
	counter(i - 1)
}

func factorial(i int) int {
	if i == 1 {
		return 1
	}

	return i * factorial(i-1)
}

func maxSquare(height, width int) int {
	if width%height == 0 {
		return height
	}

	if width > height {
		return maxSquare(width-height, height)
	} else {
		return maxSquare(height-width, width)
	}
}

func sumNumbers(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	return nums[i] + sumNumbers(nums, i+1)
}

func maxNumber(nums []int, i int) int {
	if i == len(nums)-1 {
		return nums[len(nums)-1]
	}

	if nums[i] > maxNumber(nums, i+1) {
		return nums[i]
	} else {
		return maxNumber(nums, i+1)
	}
}

func binarySearchRecursive(arr []int, x int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if arr[mid] == x {
		return mid
	} else if x < arr[mid] {
		return binarySearchRecursive(arr, x, left, mid-1)
	} else {
		return binarySearchRecursive(arr, x, mid+1, right)
	}
}

func quickSort(iArr []int) []int {
	if len(iArr) < 2 {
		return iArr
	}
	// Выбираем число для сравнения с остальными
	// Если числа в сортируемом массиве могут повторяться, лучше использовать pivot массив
	// Впрочем, сейчас переделаю )
	pivot := make([]int, 0, 2)
	pivot = append(pivot, iArr[0])
	less := make([]int, 0, len(iArr)/2)
	greater := make([]int, 0, len(iArr)/2)
	// Раскладываем значения по соответствующим спискам
	for idx, val := range iArr {
		// Условие я думаю тут ни к чему, лучше вместо range использовать i:=1, но да ладно.
		// Впрочем, для O(n log n) нужно использовать случайный элемент массива, так что i:=1 не прокатит
		if idx == 0 {
			continue
		}
		if val < pivot[0] {
			less = append(less, val)
		} else if val > pivot[0] {
			greater = append(greater, val)
		} else {
			pivot = append(pivot, val)
		}
	}
	// Го явно не про решения в одну строчку
	// Либо я туповат
	result := make([]int, 0, len(iArr))
	result = append(result, quickSort(less)...)
	result = append(result, pivot...)
	result = append(result, quickSort(greater)...)
	return result
}
