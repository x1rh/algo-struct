package binarySearch

import (
	"binarySearch/closeInterval"
	"binarySearch/closeOpenInterval"
	"binarySearch/openInterval"
	"sort"
	"testing"
)

type testCase struct {
	a             []int // 目标数组
	x             int   // 要查找的数
	lowerBoundRes int   // LowerBound()应该返回的正确结果
	upperBoundRes int   // UpperBound()应该返回的正确结果
}

var cases []testCase = []testCase{
	{[]int{0, 1, 2, 3, 4, 5, 6}, 1, 1, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6}, -1, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 6, 6, 7},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 7, 7, 7},
	{[]int{0, 1, 1, 1, 2, 3, 4, 5, 6}, 1, 1, 4},
	{[]int{0, 1, 1, 1, 1, 2, 3, 4, 5, 6}, 1, 1, 5},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 9},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 10},
}

func TestAll(t *testing.T) {
	for _, c := range cases {
		sort.Ints(c.a)
		if res := closeInterval.UpperBound(c.a, c.x); res != c.upperBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.upperBoundRes)
		}
		if res := closeInterval.LowerBound(c.a, c.x); res != c.lowerBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.lowerBoundRes)
		}

		if res := closeOpenInterval.UpperBound(c.a, c.x); res != c.upperBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.upperBoundRes)
		}
		if res := closeOpenInterval.LowerBound(c.a, c.x); res != c.lowerBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.lowerBoundRes)
		}

		if res := openInterval.UpperBound(c.a, c.x); res != c.upperBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.upperBoundRes)
		}
		if res := openInterval.LowerBound(c.a, c.x); res != c.lowerBoundRes {
			t.Fatalf("%v, res=%v, expect=%v\n", c.a, res, c.lowerBoundRes)
		}
	}
}
