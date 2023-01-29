package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("any size", func(t *testing.T){
		numbers := []int{1,2,3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	arr2 := []int{1,2,3}
	t.Run("one slice", func(t *testing.T) {
		got := SumAll(arr1)
		want := 6
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("two slices", func(t *testing.T){
		got := SumAll(arr1, arr2)
		want := 21
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

