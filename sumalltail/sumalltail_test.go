package main

import "testing"

func assertSliceEqual(t *testing.T, got []int, want []int) {
	t.Helper()
	lg := len(got)
	lw := len(want)
	if lg != lw {
		t.Errorf("len(got): %d len(want): %d\n", lg, lw)
	}
	for i := 0; i < lw; i++ {
		g := got[i]
		w := want[i]
		if w != g {
			t.Errorf("index: %d got[]: %d want: %d\n", i, g, w)
		}
	}
}

func TestSumAll(t *testing.T) {
	t.Run("non-empty slice test", func(t *testing.T) {
		got := sumall([]int{1, 3, 5}, []int{2, 4})
		want := []int{8, 4}
		assertSliceEqual(t, got, want)
	})
	t.Run("empty slice test", func(t *testing.T) {
		got := sumall([]int{}, []int{2, 4})
		want := []int{0, 4}
		assertSliceEqual(t, got, want)
	})
}
