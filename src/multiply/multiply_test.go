package multiply

import (
	"testing"
	"reflect"
)

func TestSumValues(t *testing.T) {
	records := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	expected := 24

	result := MultiplyValues(records)
	if(!reflect.DeepEqual(result, expected)){
		t.Errorf("Should be %d and it is %d:", expected, result)
	}

}
