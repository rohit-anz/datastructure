package merge

import "fmt"

func MergeSort(arr [9]int) {
	var p int
	for p=2; p <= len(arr); p=p*2 {
		for i:=0; i+p-1<len(arr); i=i+p {
			l := i
			h := i + p - 1

			mid := (l + h)/2
			arr = merge(arr, l, mid, h)
		}
	}

	if p/2 < len(arr) {
		arr = merge(arr, 0, p/2, len(arr)-1)
	}

	fmt.Println(arr)
}



func merge(arr [9]int, l int, mid int, h int) [9]int {
	i:=l
	j:= mid + 1
	k:=l

	var c  [9]int

	for i < mid && j < h {
		if arr[i] <= arr[j] {
			c[k] = arr[i]
			i = i + 1
			k = k + 1
		} else {
			c[k] = arr[j]
			j = j + 1
			k = k + 1
		}
	}

	for i < mid {
		c[k] = arr[i]
		i = i + 1
		k = k + 1
	}

	for j < h {
		c[k]= arr[j]
		j = j + 1
		k = k + 1
	}

	for i=0; i<h; i++ {
		arr[i] = c[i]
	}
	return  arr
}