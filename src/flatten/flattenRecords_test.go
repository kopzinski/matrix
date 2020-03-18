package flatten

import (
	"testing"
	"reflect"
)

func TestFlattenRecords(t *testing.T) {
	records := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	flatten := FlattenRecords(records)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if(!reflect.DeepEqual(flatten, expected)){
		t.Errorf("Should be flatten like %d and it is %d:", expected, flatten)
	}
}
