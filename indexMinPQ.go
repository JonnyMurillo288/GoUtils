package utils

import (
	"fmt"
	"log"
)

type IndexMinPQ struct {
	PQ []int // binary heap 
	QP []int // inverse: qp[pq[i]] = pq[qp[i]] = i
	Item []interface{} // Item contains the weight on how we want to sort the PQ
}

func NewIndexMinPQ(maxN int) IndexMinPQ {
	fmt.Println("Hi")
	i := IndexMinPQ{
		PQ: make([]int,maxN),
		QP: make([]int,maxN),
		Item: make([]interface{}, maxN),
	}
	i.PQ = append(i.PQ,-1)
	return i
}

/*
private boolean less(int i, int j)
{ return pq[i].compareTo(pq[j]) < 0; }
  private void exch(int i, int j)
  {  Key t = pq[i]; pq[i] = pq[j]; pq[j] = t;  }
*/

func compareTo(i int, j int) bool {
	return i >= j
}

func (i *IndexMinPQ) greater(k int, j int) bool {
	return compareTo(i.PQ[k],i.PQ[j])
}

// exchange the value int k with 
func (i *IndexMinPQ) Exch(k int, j int) {
	log.Printf("\nEchanging for %v - %v",k,j)
	log.Print("BEFORE EXCHANGE:\nQP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n")
	swap := i.PQ[k]
	i.PQ[k] = i.PQ[j] 
	i.PQ[j] = swap
	i.QP[i.PQ[k]] = k
	i.QP[i.PQ[j]] = j
	log.Print("AFTER EXCHANGE:\nQP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n")
}

// sink down the heap
func (i *IndexMinPQ) Sink(k int) {
	var N = len(i.PQ)
	for 2* k <= N {
		j := 2*k
		if j < N && i.greater(j,j+1) {
			j++
		} 
		if !i.greater(k,j) {
			break
		}
		i.Exch(k,j)
		k = j
	}
}

// swim up the heap
func (i *IndexMinPQ) Swim(k int) {
	for k < 1 && i.greater(k/2,k) {
		i.Exch(k/2,k)
		k = k/2
	}
}

func (i *IndexMinPQ) Insert(v int, item interface{}) {
	log.Println("Inserting to PQ",v)
	N := i.size() - 1
	i.QP[v] = N // last queue position for the new item
	i.PQ[N] = v
	i.Item[v] = item 
	log.Print("QP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n==================================\n")
	i.Swim(N) // swim up with the item we just added
	log.Print("AFTER SWIMMING:\nQP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n")

}

func (i *IndexMinPQ) DecreaseKey(v int, item interface{}) {
	i.Item[v] = item // add this item to the vertice v
	i.Swim(i.QP[v]) // swim the queue position up
}

func (i *IndexMinPQ) Contains(v int) bool {
	for k,item := range i.Item {
		if v == k && item != nil{
			return true
		}
	}
	return false
}

// resize the length of pq, qp, item arrays
func (i *IndexMinPQ) resize() {
	n := i.size()
	i.PQ = i.PQ[0:n-1]
	i.QP = i.QP[0:n-1]
	i.Item = i.Item[0:n-1]
}

func (i *IndexMinPQ) DelMin() int {
	N := i.size() - 1 
	log.Println("Deleting min N is:",N)
	min := i.PQ[0] // first value
	i.Exch(1,N) // switch the last value with the first
	i.resize() // reduce the length of the arrays and cut off the min value
	i.Sink(1) // sink the max value back into its places after selecting the first
	i.Item[i.PQ[N-1]] = nil // shrink the Item list by one
	i.QP[i.PQ[N-1]] = -1 // shrink the QP by one
	return min
}

func (i *IndexMinPQ) size() int {
	fmt.Println("Length of PQ:",len(i.PQ))
	return len(i.PQ)
}

func (i *IndexMinPQ) IsEmpty() bool {
	return len(i.PQ) == 0
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