package godskeqing

import (
	"fmt"
	"strings"
	"sync"
)

type Set[K comparable] map[K]struct{}

func NewSet[K comparable]() Set[K] {
	return make(Set[K])
}

func (set Set[K]) Empty() bool {
	return len(set) == 0
}

func (set Set[K]) Contains(k K) bool {
	if _, ok := set[k]; ok {
		return true
	}

	return false
}

func (set Set[K]) ContainsAll(ks []K) bool {
	for _, k := range ks {
		if !set.Contains(k) {
			return false
		}
	}

	return true
}

func (set Set[K]) Add(k K) {
	set[k] = struct{}{}
}

func (set Set[K]) AddAll(ks []K) {
	for _, k := range ks {
		set.Add(k)
	}
}

func (set Set[K]) Remove(k K) {
	delete(set, k)
}

func (set Set[K]) Merge(s Set[K]) {
	for e := range s {
		set.Add(e)
	}
}

func (set Set[K]) ToSlice() []K {
	res := make([]K, 0, len(set))
	for k := range set {
		res = append(res, k)
	}

	return res
}

func FromSlice[K comparable](slc []K) Set[K] {
	s := make(Set[K])
	for _, e := range slc {
		s.Add(e)
	}
	return s
}

func (set Set[K]) String() string {
	strbuf := new(strings.Builder)
	strbuf.WriteString("(")
	for k := range set {
		strbuf.WriteString(fmt.Sprintf("%v, ", k))
	}
	if strbuf.Len() != 1 {
		strbuf.WriteString("\b\b")
	}
	strbuf.WriteString(")")
	return strbuf.String()
}

type ConcurrentSet[K comparable] struct {
	m  Set[K]
	rw *sync.RWMutex
}

func NewConcurrentSet[K comparable]() ConcurrentSet[K] {
	return ConcurrentSet[K]{
		m:  NewSet[K](),
		rw: &sync.RWMutex{},
	}
}

func (set *ConcurrentSet[K]) Empty() bool {
	set.rw.RLock()
	defer set.rw.RUnlock()

	return set.m.Empty()
}

func (set *ConcurrentSet[K]) Contains(k K) bool {
	set.rw.RLock()
	defer set.rw.RUnlock()

	return set.m.Contains(k)
}

func (set *ConcurrentSet[K]) ContainsAll(ks []K) bool {
	set.rw.RLock()
	defer set.rw.RUnlock()

	return set.m.ContainsAll(ks)
}

func (set *ConcurrentSet[K]) Add(k K) {
	set.rw.Lock()
	defer set.rw.Unlock()

	set.m.Add(k)
}

func (set *ConcurrentSet[K]) AddAll(ks []K) {
	set.rw.Lock()
	defer set.rw.Unlock()

	set.m.AddAll(ks)
}

func (set *ConcurrentSet[K]) Remove(k K) {
	set.rw.Lock()
	defer set.rw.Unlock()

	set.m.Remove(k)
}

func (set *ConcurrentSet[K]) Merge(s Set[K]) {
	set.rw.Lock()
	defer set.rw.Unlock()

	set.m.Merge(s)
}

func (set *ConcurrentSet[K]) ToSlice() []K {
	set.rw.RLock()
	defer set.rw.RUnlock()

	return set.m.ToSlice()
}

func ConcurrentFromSlice[K comparable](slc []K) ConcurrentSet[K] {
	s := NewConcurrentSet[K]()
	for _, e := range slc {
		s.Add(e)
	}
	return s
}
