package internal

import (
	"cmp"
	"iter"
)

// specialized version (for better performance)
type Heap[T cmp.Ordered] []T

func NewHeap[T cmp.Ordered](vs []T) Heap[T] {
	h := Heap[T](vs)
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- { // n/2-1: last (in depth order) parent
		h.down(i, n)
	}
	return h
}

// Note: in order to prevent a func literal from escaping to the heap,
// we deliberately design this method as an iter.Seq[T] rather than
// as a iter.Seq[T] factory.
func (h Heap[T]) Iterator(yield func(T) bool) {
	var v T
	for range len(h) {
		v, h = h.pop()
		if !yield(v) {
			break
		}
	}
}

var _ iter.Seq[int] = Heap[int]{}.Iterator // compile-time check

func (h Heap[_]) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap[_]) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Note: this implementation gets rid of one level of indirection compared to
// container/heap's implementation.
func (h Heap[T]) pop() (T, Heap[T]) {
	n := len(h) - 1
	h.swap(0, n)
	h.down(0, n)
	x := h[n]
	h = h[0:n] // TODO: should the element of index n in h.s be cleared?
	return x, h
}

func (h Heap[_]) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
}
