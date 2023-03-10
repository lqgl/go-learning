package main

import (
	"reflect"
	"testing"
)

// go test -cover

func TestSum(t *testing.T) {

	// t.Run("collections of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("collectons of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// reflect.DeepEqual 不是「类型安全」的
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
