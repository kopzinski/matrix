package converter

import (
	"testing"
	"reflect"
)

func TestConvertStringMatrixToInt(t *testing.T) {
	stringMatrix := [][]string{
		[]string{"1", "2"},
		[]string{"3", "4"}}

	expected := [][]int{
		[]int{1, 2},
		[]int{3, 4}}

	intMatrix, _ := ConvertStringMatrixToInt(stringMatrix)

	if(!reflect.DeepEqual(intMatrix, expected)){
		t.Errorf("Should be converted into int matrix like %d and it is %d:", expected, intMatrix)
	}
}

func TestConvertIntArrayToString(t *testing.T) {
	intArray := []int{1, 2}
	expected := []string{"1", "2"}

	stringArray := ConvertIntArrayToString(intArray)

	if(!reflect.DeepEqual(stringArray, expected)) {
		t.Errorf("Should be converted into string array like %s and it is %s:", expected, stringArray)
	}
}
