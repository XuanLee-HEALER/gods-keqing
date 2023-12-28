package godskeqing_test

import (
	"sync"
	"testing"

	godskeqing "github.com/XuanLee-HEALER/gods-keqing"
)

func TestQueue(t *testing.T) {
	q := godskeqing.NewQueue()

	ori := []any{1, "x", 2, "y", 3, "z"}
	des := ori

	for _, v := range ori {
		q.Enqueue(v)
	}

	counter := 0
	for !q.Empty() {
		// t.Log(counter)
		if v := q.Dequeue(); v == nil || v != des[counter] {
			t.FailNow()
		}
		counter++
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := godskeqing.NewQueue()
	if q.Dequeue() != nil {
		t.FailNow()
	}
}

func TestConcurrentQueueDequeueEmpty(t *testing.T) {
	q := godskeqing.NewConcurrentQueue()
	if q.Dequeue() != nil {
		t.FailNow()
	}
}

func TestConcurrentQueue(t *testing.T) {
	q := godskeqing.NewConcurrentQueue()

	ori := []any{1, "x", 2, "y", 3, "z"}

	wg := new(sync.WaitGroup)
	for i := range ori {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			q.Enqueue(ori[idx])
		}(i)
	}

	wg.Wait()

	for !q.Empty() {
		t.Log(q.Dequeue())
	}
}

func TestConcurrentPriorityQueueDequeueEmpty(t *testing.T) {
	q := godskeqing.NewConcurrentPriorityQueue(nil)
	if q.Dequeue() != nil {
		t.FailNow()
	}
}

func TestConcurrentPriorityQueue(t *testing.T) {
	q := godskeqing.NewConcurrentPriorityQueue(func(a, b interface{}) int {
		av := a.(int)
		bv := b.(int)
		if av < bv {
			return -1
		} else if av == bv {
			return 0
		} else {
			return 1
		}
	})

	ori := []any{1, 2, 1, 3, 4, 1, 2, 3, 4, 5, 2, 2}

	wg := new(sync.WaitGroup)
	for i := range ori {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			q.Enqueue(ori[idx])
		}(i)
	}

	wg.Wait()

	for !q.Empty() {
		t.Log(q.Dequeue())
	}
}

func TestStack(t *testing.T) {
	s := godskeqing.NewStack()

	s.Push(1)
	s.Push("something")
	s.Push(2)
	s.Push(3)
	s.Push("anything")

	for !s.Empty() {
		t.Logf("value: %v", s.Remove())
	}
}
