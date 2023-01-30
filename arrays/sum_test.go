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

func TestSumAllTails(t *testing.T) {
	arr1 := []int{1,2,3,4,5}
	arr2 := []int{1,2,3}

	checkSums := func(t testing.TB, got, want [] int) {
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("one slice", func(t *testing.T) {
		got := SumAllTails(arr2)
		want := []int{5}
		checkSums(t, got, want)
	})
	t.Run("two slices", func(t *testing.T){
		got := SumAllTails(arr1, arr2)
		want := []int{14, 5}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails(arr2, []int{})
		want := []int{5, 0}
		checkSums(t, got, want)
	})

}

