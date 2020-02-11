package count

import "fmt"

func Sort(arr []int) {
	max := findMaxElement(arr)
	var auxilaryArray [100]int
	var resultArray []int

	for i:= 0; i <= max; i++ {
		auxilaryArray[i] = 0
	}

	for i:= 0; i < len(arr); i++ {
		auxilaryArray[arr[i]] =  auxilaryArray[arr[i]] + 1
	}

	for i:= 0; i <= max; i++ {
		for auxilaryArray[i] > 0 {
			resultArray = append(resultArray, i)
			auxilaryArray[i] = auxilaryArray[i] - 1
		}
	}
	fmt.Println(resultArray)
}

func findMaxElement(arr []int) int {
	max := arr[0]

	for _, element := range arr {
		if element > max {
			max = element
		}
	}

	return max
}