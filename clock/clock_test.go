// Copyright 2016 Markus Sonderegger. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clock

import "testing"

func TestUpdateAndSearch(t *testing.T) {
	vc := New(nil)
	vc.Update("id6", 0)
	vc.Update("id1", 0)
	vc.Update("id3", 0)
	vc.Update("id5", 0)
	vc.Update("id2", 0)
	vc.Update("id4", 0)

	check := func(got, want int, found, must bool) {
		if found != must {
			t.Fatalf("update: unexpcted item slice state, want %v, got %v", must, found)
		}
		if got != want {
			t.Fatalf("update: expected item at index %d, got %d", want, got)
		}
	}

	i, found := vc.find("id3")
	check(i, 2, found, true)
	i, found = vc.find("id1")
	check(i, 0, found, true)
	i, found = vc.find("id6")
	check(i, 5, found, true)
	i, found = vc.find("id9")
	check(i, 6, found, false)
}

func TestUpdateAndCompare(t *testing.T) {
	vc1, vc2 := New(nil), New(nil)
	check := func(vc1, vc2 *Clock, cond Condition, want bool) {
		if Compare(vc1, vc2, cond) != want {
			t.Fatalf("compare: unecpected compare result %v", !want)
		}
	}

	check(vc1, vc2, Equals, true)
	check(vc2, vc1, Equals, true)
	check(vc1, vc2, ^Equals, false)
	check(vc2, vc1, ^Equals, false)

	vc2.Update("idA", 0)
	check(vc1, vc2, Descendant, true)
	check(vc1, vc2, Equals, false)
	check(vc1, vc2, ^Descendant, false)
	check(vc2, vc1, Ancestor, true)
	check(vc2, vc1, Equals, false)
	check(vc2, vc1, ^Ancestor, false)

	vc1.Update("idA", 0)
	check(vc1, vc2, Equals, true)
	check(vc2, vc1, Equals, true)
	check(vc1, vc2, ^Equals, false)
	check(vc2, vc1, ^Equals, false)

	vc1.Update("idB", 0)
	check(vc1, vc2, Ancestor, true)
	check(vc1, vc2, ^Ancestor, false)
	check(vc2, vc1, Descendant, true)
	check(vc2, vc1, ^Descendant, false)

	vc2.Update("idA", 0)
	check(vc1, vc2, Concurrent, true)
	check(vc1, vc2, ^Concurrent, false)
	check(vc2, vc1, Concurrent, true)
	check(vc2, vc1, ^Concurrent, false)

	vc2.Update("idB", 0)
	check(vc1, vc2, Descendant, true)
	check(vc1, vc2, ^Descendant, false)
	check(vc2, vc1, Ancestor, true)
	check(vc2, vc1, ^Ancestor, false)

	vc2.Update("idB", 0)
	check(vc1, vc2, Descendant, true)
	check(vc1, vc2, ^Descendant, false)
	check(vc2, vc1, Ancestor, true)
	check(vc2, vc1, ^Ancestor, false)
}

func TestCompareWithMissing(t *testing.T) {
	vc1, vc2 := New(nil), New(nil)
	check := func(vc1, vc2 *Clock, cond Condition, want bool) {
		if Compare(vc1, vc2, cond) != want {
			t.Fatalf("compare: unecpected compare result %v", !want)
		}
	}

	vc1.Update("idA", 0)
	vc1.Update("idC", 0)
	vc2.Update("idB", 0)

	check(vc1, vc2, Equals, false)
	check(vc2, vc1, Equals, false)
	check(vc2, vc1, Descendant, false)
	check(vc2, vc1, Ancestor, false)
	check(vc2, vc1, Concurrent, true)

	vc1.Update("idD", 0)

	check(vc1, vc2, Equals, false)
	check(vc2, vc1, Equals, false)
	check(vc2, vc1, Descendant, false)
	check(vc2, vc1, Ancestor, false)
	check(vc2, vc1, Concurrent, true)
}

func TestMerge(t *testing.T) {
	check := func(vc1, vc2 *Clock, cond Condition, want bool) {
		if Compare(vc1, vc2, cond) != want {
			t.Fatalf("compare: unecpected compare result %v", !want)
		}
	}

	vc1, vc2 := New(nil), New(nil)
	vc1.Update("idA", 0)
	vc1.Update("idA", 0) // counter > than vc2
	vc1.Update("idB", 0) // counter < than vc2
	vc1.Update("idC", 0) // counter not in vc2
	vc1.Update("idD", 0) // counter not in vc2

	vc2.Update("idA", 0)
	vc2.Update("idB", 0)
	vc2.Update("idB", 0)
	vc2.Update("idE", 0) // counter not in vc1
	vc2.Update("idF", 0) // counter not in vc1

	vc2.Merge(vc1)

	check(vc1, vc2, Descendant, true)
	check(vc1, vc2, ^Descendant, false)
	check(vc2, vc1, Ancestor, true)
	check(vc2, vc1, ^Ancestor, false)

	vc1.Update("idB", 0)
	vc1.Update("idE", 0)
	vc1.Update("idF", 0)

	check(vc2, vc1, Equals, true)
}
