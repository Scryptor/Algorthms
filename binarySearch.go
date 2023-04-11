package main

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		// Делим пополам, чтобы дальше сравнить половинки с item
		mid := (low + high) / 2
		guess := list[mid]
		// проверяем не попали ли мы сразу
		if guess == item {
			return mid
		}
		// если число из середины промежутка > нужного
		// прописываем новый дипазон для High, то бишь понижаем промежуток в меньшую сторону
		if guess > item {
			high = mid - 1
			// если меньше нужного, то повышаем нижнюю границу
		} else {
			low = mid + 1
		}
		// И гоняем пока не найдем
	}
	return -1
}
