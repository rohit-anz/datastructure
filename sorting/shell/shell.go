package shell

import "fmt"

func Sort(arr [8]int) {
	gap := (len(arr)-1)/2
	if len(arr) <= 1 {
		fmt.Println("array does not have sufficient elements")
	} else {
		for i:= gap; i>0; i = i/2 {
			for j:=0; j+i < len(arr); j++ {
				if arr[j] > arr[j + i] {
					temp:= arr[j]
					arr[j] = arr[j+i]
					arr[j+i] = temp
				}

				tempGap:= j - i

				for tempGap >= 0 {
					if arr[j] <= arr[tempGap] {
						temp:= arr[j]
						arr[j] = arr[tempGap]
						arr[tempGap] = temp
						tempGap = tempGap - i
					} else {
						break
					}
				}
			}
		}
 	}
	fmt.Println(arr)
}