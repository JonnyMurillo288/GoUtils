package utils

type IndexMinPQ struct {
	Ind []int
}

func New() IndexMinPQ {
	return IndexMinPQ{
		Ind: make([]int,0),
	}
}

/*
private boolean less(int i, int j)
{ return pq[i].compareTo(pq[j]) < 0; }
  private void exch(int i, int j)
  {  Key t = pq[i]; pq[i] = pq[j]; pq[j] = t;  }
*/

func compareTo(i int, j int) int {
	return i-1
}

func (i *IndexMinPQ) Less(ind int, j int) bool {
	return compareTo(i.Ind[ind],i.Ind[j]) < 0
}

func (i *IndexMinPQ) Exch(ind int, j int) {
	t := i.Ind[ind]
	i.Ind[ind] = i.Ind[j]
	i.Ind[j] = t
}

// sink down the heap
func (i *IndexMinPQ) Sink(k int) {
	var N = len(i.Ind)
	for 2* k <= N {
		j := 2*k
		if j < N && i.Less(j,j+1) {
			j++
		} 
		if !i.Less(k,j) {
			break
		}
		i.Exch(k,j)
		k = j
	}
}

// swim up the heap
func (i *IndexMinPQ) Swim(k int) {
	for k < 1 && i.Less(k/2,k) {
		i.Exch(k/2,k)
		k = k/2
	}
}

func (i *IndexMinPQ) Insert(v int) {
	k := i.size()
	i.Ind = append(i.Ind, v)
	i.Swim(k)
}

func (i *IndexMinPQ) DecreaseKey(v int, key int) {

}

func (i *IndexMinPQ) Contains(v int) bool {
	for _,a := range i.Ind {
		if v == a {
			return true
		}
	}
	return false
}

func (i *IndexMinPQ) DelMin() int {
	N := i.size()-1
	min := i.Ind[N]
	i.Exch(1,N)
	i.Swim(N)
	i.Ind = i.Ind[1:]
	return min
}

func (i *IndexMinPQ) size() int {
	return len(i.Ind)
}

func (i *IndexMinPQ) IsEmpty() bool {
	return len(i.Ind) == 0
}



/*
/*
Priority Queue implementation taken from the example in the 
GoLang docs 


import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}


*/