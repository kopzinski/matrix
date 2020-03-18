package invert

import (
	"testing"
	"reflect"

)

func TestInvertRecords(t *testing.T) {
	records := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	inverted := InvertRecords(records)

	expected := [][]int{
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
	}
	if(!reflect.DeepEqual(inverted, expected)){
		t.Errorf("Should be inverted like %d and it is %d:", expected, inverted)
	}

}
