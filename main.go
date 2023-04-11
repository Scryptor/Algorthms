package main

import (
	"fmt"
)

func main() {
	//graphs
	queue := NewQueue()

	igor := Person{
		"Igor",
		[]Person{
			{
				"Vova",
				[]Person{},
			},
			{
				"Sasha",
				[]Person{},
			},
		},
	}

	maxim := Person{
		"Maxim",
		[]Person{
			{
				"Ivan",
				[]Person{},
			},
			{
				"Sergej",
				[]Person{},
			},
			{
				"Sasha",
				[]Person{},
			},
		},
	}

	dima := Person{
		"Dmitriy",
		[]Person{
			{
				"Sergej",
				[]Person{},
			},
			{
				"Igor",
				[]Person{},
			},
		},
	}
	artem := Person{
		"Dmitriy",
		[]Person{
			{
				"Petr",
				[]Person{},
			},
		},
	}
	vasya := Person{
		"Vasya",
		[]Person{},
	}
	queue.Enqueue(igor, maxim, dima, vasya, artem)
	count := 0
	for !queue.IsEmpty() {
		count++
		p, err := queue.Dequeue()
		if err != nil {
			panic(err)
		}
		if _, ok := queue.Checked[p.Name]; ok {
			continue
		} else {
			queue.Checked[p.Name] = struct{}{}
		}
		if p.Name != "Ivan" {
			queue.Enqueue(p.Friends...)
		} else {
			fmt.Println("Ivan found in", count, "steps")
			break
		}
	}
	//recursion
	fmt.Print("Count from 7: ")
	counter(7)
	fmt.Println()
	fmt.Println("3 factorial:", factorial(3))
	fmt.Println("MaxSquare in 1680 and 640:", maxSquare(1680, 640))
	fmt.Println("Sum of numbers 2,4,6 is: ", sumNumbers([]int{2, 4, 6}, 0))
	fmt.Println("Max number 2,6,1,8,3,5:", maxNumber([]int{2, 6, 1, 8, 3, 5}, 0))
	fmt.Println("Index of 7 recursive is:", binarySearchRecursive([]int{0, 1, 2, 3, 5, 7, 9, 10, 15, 55}, 7, 0, 9))
	fmt.Println("Quick sort recursive 10,5,2,3,3,3,3:", quickSort([]int{10, 5, 2, 3, 3, 3, 3}))
	//arraySortOn2
	listSort := []int{0, 3, 6, 8, 2, 9, 0, 0}
	fmt.Println(fmt.Sprintf("Sorted array %v:", listSort), selectionSort(listSort))

	//binarySearch
	myList := []int{1, 3, 5, 7, 9}
	fmt.Println("Index of 3 is:", binarySearch(myList, 3))

}
