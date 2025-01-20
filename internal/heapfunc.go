package internal

import "iter"

type HeapFunc[T any] struct {
	s       []T
	compare func(x, y T) int
}

func NewHeapFunc[T any](vs []T, compare func(T, T) int) HeapFunc[T] {
	h := HeapFunc[T]{
		s:       vs,
		compare: compare,
	}
	// heapify
	n := h.len()
	for i := n/2 - 1; i >= 0; i-- { // n/2-1: last (in depth order) parent
		h.down(i, n)
	}
	return h
}

// Note: in order to prevent a func literal from escaping to the heap,
// we deliberately design this method as an iter.Seq[T] rather than
// as a iter.Seq[T] factory.
func (h HeapFunc[T]) Iterator(yield func(T) bool) {
	var v T
	for range h.len() {
		v, h = h.pop()
		if !yield(v) {
			break
		}
	}
}

var _ iter.Seq[int] = HeapFunc[int]{}.Iterator // compile-time check

func (h HeapFunc[_]) less(i, j int) bool {
	return h.compare(h.s[i], h.s[j]) < 0
}

func (h HeapFunc[_]) swap(i, j int) {
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func (h HeapFunc[_]) len() int {
	return len(h.s)
}

// Note: this implementation gets rid of one level of indirection compared to
// container/heap's implementation.
func (h HeapFunc[T]) pop() (T, HeapFunc[T]) {
	n := h.len() - 1
	h.swap(0, n)
	h.down(0, n)
	x := h.s[n]
	h.s = h.s[0:n] // TODO: should the element of index n in h.s be cleared?
	return x, h
}

func (h HeapFunc[_]) down(i, n int) {
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
