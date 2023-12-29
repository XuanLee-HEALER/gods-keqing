package godskeqing_test

import (
	"sync"
	"testing"
	"time"

	godskeqing "github.com/XuanLee-HEALER/gods-keqing"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := godskeqing.NewSet[int]()
	assert.NotNil(t, s, "new set is nil")
}

func TestEmpty(t *testing.T) {
	s := godskeqing.NewSet[int]()
	assert.Equal(t, true, s.Empty(), "new set is not empty")
}

func TestAdd(t *testing.T) {
	s := godskeqing.NewSet[int]()
	s.Add(1)
	assert.Contains(t, s, 1, "1 was added unsuccessfully")
}

func TestAddAll(t *testing.T) {
	s := godskeqing.NewSet[int]()
	s.AddAll([]int{1, 2, 3})
	assert.Contains(t, s, 1, "1 was added unsuccessfully")
	assert.Contains(t, s, 2, "2 was added unsuccessfully")
	assert.Contains(t, s, 3, "3 was added unsuccessfully")
}

func TestContains(t *testing.T) {
	s := godskeqing.NewSet[int]()
	assert.Equal(t, false, s.Contains(1), "empty set contains 1")
	s.Add(1)
	assert.Equal(t, true, s.Contains(1), "1 was added unsuccessfully")
}

func TestContainsAll(t *testing.T) {
	s := godskeqing.NewSet[int]()
	assert.Equal(t, false, s.ContainsAll([]int{1, 2, 3}), "empty set contains [1,2,3]")
	s.AddAll([]int{1, 2, 3})
	assert.Equal(t, true, s.ContainsAll([]int{1, 2, 3}), "[1,2,3] was added unsuccessfully")
}

func TestRemove(t *testing.T) {
	s := godskeqing.NewSet[int]()
	s.AddAll([]int{1, 2, 3})
	s.Remove(1)
	assert.Equal(t, true, s.ContainsAll([]int{2, 3}), "1 was deleted unsuccessfully")
}

func TestMerge(t *testing.T) {
	s1 := godskeqing.NewSet[int]()
	s2 := godskeqing.NewSet[int]()
	s1.AddAll([]int{1, 2, 2, 3})
	s2.AddAll([]int{2, 3, 4, 5})
	s1.Merge(s2)
	assert.ElementsMatch(t, s1.ToSlice(), []int{1, 2, 3, 4, 5}, "failed to merge")
}

func TestToSlice(t *testing.T) {
	s := godskeqing.NewSet[int]()
	s.AddAll([]int{1, 2, 3})
	l := s.ToSlice()
	assert.ElementsMatch(t, []int{1, 2, 3}, l, "set converted to slice has changed")
	l = []int{}
	assert.ElementsMatch(t, s.ToSlice(), []int{1, 2, 3}, "set changed by new reference")
	assert.ElementsMatch(t, l, []int{}, "")
}

func TestFromSlice(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 5, 5, 6}
	s := godskeqing.FromSlice[int](arr)
	assert.ElementsMatch(t, s.ToSlice(), []int{1, 2, 3, 4, 5, 6}, "set converted by slice is wrong")
}

func TestConcurrentSet(t *testing.T) {
	s := godskeqing.NewConcurrentSet[int]()

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			for xi := 1; xi <= 5; xi++ {
				if xi*i/2 == 0 {
					s.Add(xi)
				}
			}
		}(i)
	}
	wg.Wait()

	assert.ElementsMatch(t, s.ToSlice(), []int{1, 2, 3, 4, 5}, "wrong values")
}
