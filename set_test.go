package godskeqing_test

import (
	"testing"

	godskeqing "github.com/XuanLee-HEALER/gods-keqing"
)

func TestSetMerge(t *testing.T) {
	s1 := make(godskeqing.Set[int])
	s2 := make(godskeqing.Set[int])
	s1.AddList([]int{1, 2, 3})
	s2.AddList([]int{2, 3, 4})
	s1.Merge(s2)
	if len(s1) != 4 {
		t.Fail()
	}
	t.Logf("merge result: %v", s1)
}
