//import "container/heap"

type KthLargest struct {
	h *MinHeap
	k int
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h *MinHeap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeap(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return KthLargest{&h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	return (*this.h)[0]
}