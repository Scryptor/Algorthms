package main

// Находит в массиве индекс наименьшего значения
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for idx, val := range arr {
		if val < smallest {
			smallest = val
			smallestIndex = idx
		}
	}
	return smallestIndex
}

// Сортирует массив
func selectionSort(arr []int) []int {
	resArr := make([]int, 0, len(arr))

	for range arr {
		smallest := findSmallest(arr)
		resArr = append(resArr, arr[smallest])
		arr = remove(arr, smallest)
	}

	return resArr
}

// Удаляет число из массива не сохраняя его позицию, что быстро и не критично в данной ситуации
// Так же возможны ошибки, ибо нет проверки крайних значений массива
func remove(s []int, i int) []int {
	// кладет в [i] последний элемент массива
	s[i] = s[len(s)-1]
	// возвращает массив без последнего элемента
	return s[:len(s)-1]
}

// Удаляет число из массива, сохраняя его позицию
func removeSlow(slice []int, s int) []int {
	// Возвращает срез  [i,i,i,X,i,i]
	return append(slice[:s], slice[s+1:]...)
}

// на больших массивах значений разница в скорости между функциями может достигать миллионы раз
