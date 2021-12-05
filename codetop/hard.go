package codetop

import (
	"container/heap"
	"sort"
)

//////////////////////////////////////////////////////////////////////////////////////////////
// leetcode No.480
// 这题不会，先琢磨透参考答案，再自己重新做
// 思路是：维护两个优先级队列，一个大顶堆，一个小顶堆，保证大顶堆的个数最多比小顶堆的个数多一个
// 确保： 大顶堆的堆顶和小顶堆的堆顶 数组区间是最中间的两个数
// 		 移除堆元素的时候维护堆的size，提前把size减掉（不影响中位数的元素可以延迟删除），
//       其他情况push的时候加，pop的时候减（这里卡了半天，忽略维护size的细节， 然后就一直调试，太坑了）
/////////////////////////////////////////////////////////////////////////////////////////////

type SmallIntHeap struct {
	sort.IntSlice
	size int
}

func (h *SmallIntHeap) Push(v interface{}) {
	x := v.(int)
	h.IntSlice = append(h.IntSlice, -x)
}

func (h *SmallIntHeap) Pop() interface{} {
	t := (*&h.IntSlice)[len(*&h.IntSlice)-1]
	h.IntSlice = (h.IntSlice)[:len(h.IntSlice)-1]
	return -t
}

func (h *SmallIntHeap) RunDeferDelete(deferDelete map[int]int) {
	for len(h.IntSlice) > 0 {
		top := -h.IntSlice[0]
		if _, ok := deferDelete[top]; ok {
			heap.Pop(h)
			deferDelete[top]--
			if deferDelete[top] == 0 {
				delete(deferDelete, top)
			}
		} else {
			break
		}
	}
}

type BigIntHeap struct {
	sort.IntSlice
	size int
}

func (h *BigIntHeap) Push(v interface{}) {
	x := v.(int)
	h.IntSlice = append(h.IntSlice, x)
}

func (h *BigIntHeap) Pop() interface{} {
	t := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return t
}

func (h *BigIntHeap) RunDeferDelete(deferDelete map[int]int) {
	for len(h.IntSlice) > 0 {
		top := h.IntSlice[0]
		if _, ok := deferDelete[top]; ok {
			heap.Pop(h)
			deferDelete[top]--
			if deferDelete[top] == 0 {
				delete(deferDelete, top)
			}
		} else {
			break
		}
	}
}

type MedianInt struct {
	deferDelete map[int]int
	small       SmallIntHeap
	large       BigIntHeap
}

func NewMedianInt() *MedianInt {
	return &MedianInt{
		small:       SmallIntHeap{},
		large:       BigIntHeap{},
		deferDelete: make(map[int]int),
	}
}

func (m *MedianInt) Insert(x int) {
	if len(m.small.IntSlice) == 0 || x <= -m.small.IntSlice[0] {
		heap.Push(&m.small, x)
		m.small.size++
	} else {
		heap.Push(&m.large, x)
		m.large.size++
	}
	m.Balance()
}

func (m *MedianInt) Delete(x int) {
	if v, ok := m.deferDelete[x]; ok {
		m.deferDelete[x] = v + 1
	} else {
		m.deferDelete[x] = 1
	}
	if x <= -m.small.IntSlice[0] {
		m.small.size--
		if x == -m.small.IntSlice[0] {
			m.small.RunDeferDelete(m.deferDelete)
		}
	} else {
		m.large.size--
		if m.large.IntSlice[0] == x {
			m.large.RunDeferDelete(m.deferDelete)
		}
	}
	m.Balance()
}

func (m *MedianInt) Balance() {
	if m.small.size-1 > m.large.size {
		top := heap.Pop(&m.small)
		m.small.size--
		heap.Push(&m.large, top)
		m.large.size++
		m.small.RunDeferDelete(m.deferDelete)
	} else if m.large.size > m.small.size {
		top := heap.Pop(&m.large)
		m.large.size--
		heap.Push(&m.small, top)
		m.small.size++
		m.large.RunDeferDelete(m.deferDelete)
	}
}

func (m *MedianInt) GetMedian(k int) float64 {
	if k&1 > 0 {
		return float64(-m.small.IntSlice[0])
	} else {
		return float64(-m.small.IntSlice[0]+m.large.IntSlice[0]) / 2
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var ans []float64
	m := NewMedianInt()
	for i := 0; i < k; i++ {
		m.Insert(nums[i])
	}
	ans = append(ans, m.GetMedian(k))
	for i := k; i < len(nums); i++ {
		m.Insert(nums[i])
		m.Delete(nums[i-k])
		ans = append(ans, m.GetMedian(k))
	}
	return ans
}
