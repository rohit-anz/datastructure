package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type expectedResult struct {
	result	[]int
}

var HeapSortTests = []struct{
	name 		string
	input		[]int
	want 		expectedResult
}{
	{
		name: 	"SortTheUnsortedArray",
		input:	[]int{2, 9, 5, 3, 6},
		want:	expectedResult{result:[]int{9, 6, 5, 3, 2}},

	}, {
		name: 	"SortArrayOfLargeNumber",
		input:	[]int{190, 2000, 150, 2100, 500, 300, 600, 250},
		want:   expectedResult{result:[]int{2100, 2000, 600, 500, 300, 250, 190, 150},},
	},
}

func TestHeapSort(t *testing.T) {
	for _, tt := range HeapSortTests {
		tt:= tt
		t.Run(tt.name, func(t *testing.T) {
			actualResult := HeapSort(tt.input)
			assert.Equal(t, tt.want.result, actualResult)
		})
	}
}