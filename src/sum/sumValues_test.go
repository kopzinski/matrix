package sum

import (
	"reflect"
	"testing"
)

func TestSumValues(t *testing.T) {
	records := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	expected := 10

	result := SumValues(records)
	if(!reflect.DeepEqual(result, expected)){
		t.Errorf("Should be %d and it is %d:", expected, result)
	}
	

}
