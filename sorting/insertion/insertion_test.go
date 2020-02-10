package insertion

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type result struct {
	result	[]int
}

var InsertionSortTests = []struct{
	name 		string
	input		[]int
	want 		result
}{
	{
		name: 	"InsertionSortTheUnsortedArray",
		input:	[]int{2, 9, 5, 3, 6},
		want:	result{result:[]int{2, 3, 5, 6, 9}},
	}, {
		name: 	"InsertionSortArrayOfLargeNumber",
		input:	[]int{190, 2000, 150, 2100, 500, 300, 600, 250},
		want:   result{result:[]int{150, 190, 250, 300, 500, 600, 2000, 2100},},
	},
}

func TestInserionSort(t *testing.T) {
	for _, tt := range InsertionSortTests {
		tt:= tt
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.input)
			assert.Equal(t, tt.want.result, tt.input)
		})
	}
}