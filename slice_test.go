package gutil

import "testing"

func TestSliceEqualLoosely(t *testing.T) {
	if !SliceEqualLoosely([]int{1, 2, 3}, []int{3, 2, 1}) {
		t.Fatal()
	}
	if SliceEqualLoosely([]int{1, 2}, []int{3, 2, 1}) {
		t.Fatal()
	}
	if !SliceEqualLoosely([]int{1, 3, 2, 2}, []int{3, 2, 1, 2}) {
		t.Fatal()
	}
	if SliceEqualLoosely([]int{1, 3, 2}, []int{3, 2, 1, 2}) {
		t.Fatal()
	}
	if !SliceEqualLoosely([]int{}, []int{}) {
		t.Fatal()
	}
}

func TestSliceEqual(t *testing.T) {
	if !SliceEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fatal()
	}
	if !SliceEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fatal()
	}
	if !SliceEqual([]int{}, []int{}) {
		t.Fatal()
	}
	if SliceEqual([]int{1}, []int{}) {
		t.Fatal()
	}
	if SliceEqual([]int{1, 2}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func TestSliceUnion(t *testing.T) {
	assertTrue := func(a, b, res []int) {
		c := SliceUnion(a, b)
		if !SliceEqualLoosely(c, res) {
			t.Fatal(a, b, "expected:", res, "got:", c)
		}
	}
	assertTrue([]int{}, []int{}, []int{})
	assertTrue([]int{}, []int{1}, []int{1})
	assertTrue([]int{}, []int{1, 1}, []int{1})
	assertTrue([]int{1}, []int{2}, []int{1, 2})
	assertTrue([]int{1, 2}, []int{2}, []int{1, 2})
	assertTrue([]int{1, 2, 2}, []int{2}, []int{1, 2})
	assertTrue([]int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2, 3})
}

func TestSliceIntersection(t *testing.T) {
	assertTrue := func(a, b, res []int) {
		c := SliceIntersection(a, b)
		if !SliceEqualLoosely(c, res) {
			t.Fatal(a, b, "expected:", res, "got:", c)
		}
	}
	assertTrue([]int{}, []int{}, []int{})
	assertTrue([]int{}, []int{1}, []int{})
	assertTrue([]int{1, 2}, []int{1}, []int{1})
	assertTrue([]int{1, 2}, []int{1, 2, 3}, []int{1, 2})
	assertTrue([]int{1, 2, 2, 3, 4}, []int{1, 1, 2, 3}, []int{1, 2, 3})
}
