package heap

import "fmt"

func CreateHeap(arr []int) {
	for index, _ := range arr {
		if index > 1 {
			insert(arr, index)
		}
	}
}

func insert(arr []int, index int) {
	for index/2 >= 1 {
		parent := index/2
		if arr[index] > arr[parent] {
			temp := arr[index]
			arr[index] = arr[parent]
			arr[parent] = temp
		}
		index = index/2
	}
}

func DeleteElement(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		fmt.Println("No element available in heap")
		return []int{}, -1
	} else {
		element := arr[1]
		arr[1] = arr[len(arr)-1]
		arr = arr[: len(arr)-1]
		index := 1

		for index*2 <= len(arr)-1 {
			leftChild := index * 2
			rightChild := index * 2 + 1

			resultIndex := rightChild
			if rightChild > len(arr) - 1 || arr[leftChild] > arr[rightChild] {
				resultIndex = leftChild
			}

			if arr[resultIndex] > arr[index] {
				temp := arr[resultIndex]
				arr[resultIndex] = arr[index]
				arr[index] = temp
				index = resultIndex
			}
			index = index * 2
		}
		return arr, element
	}
}

func HeapSort(arr []int) []int {
	heapArray := []int{0}
	heapArray = append(heapArray, arr...)
	CreateHeap(heapArray)

	arr = heapArray
	sortedArray:= []int{}
	for index, _ := range arr {
		if index > 0 {
			sortedArray = append(sortedArray, arr[1])
			arr[1] = arr[len(arr)-1]
			arr = arr[: len(arr)-1]

			rootIndex := 1
			for rootIndex * 2 < len(arr) {
				leftChild := 2*rootIndex
				rightChild := leftChild + 1

				resultIndex := rightChild

				if rightChild > len(arr) - 1 || arr[leftChild] > arr[rightChild] {
					resultIndex = leftChild
				}

				if arr[resultIndex] > arr[rootIndex] {
					temp := arr[resultIndex]
					arr[resultIndex] = arr[rootIndex]
					arr[rootIndex] = temp
				}
				rootIndex = rootIndex * 2
			}
		}
	}
	return sortedArray
}