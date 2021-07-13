package utils

import (
	"fmt"
	"log"
)

type IndexMinPQ struct {
	PQ []int // binary heap 
	QP []int // inverse: qp[pq[i]] = pq[qp[i]] = i
	Item []float64 // Item contains the weight on how we want to sort the PQ
}

func NewIndexMinPQ(maxN int) IndexMinPQ {
	fmt.Println("Hi")
	i := IndexMinPQ{
		PQ: make([]int,maxN+1),
		QP: make([]int,maxN+1),
		Item: make([]float64, maxN+1),
	}
	for j := 0; j < len(i.PQ); j++ {
		i.QP[j] = -1
	}
	return i
}

/*
private boolean less(int i, int j)
{ return pq[i].compareTo(pq[j]) < 0; }
  private void exch(int i, int j)
  {  Key t = pq[i]; pq[i] = pq[j]; pq[j] = t;  }
*/

func compareTo(i float64, j float64) bool {
	return i > j
}

func (i *IndexMinPQ) greater(k int, j int) bool {
	return compareTo(i.Item[k],i.Item[j])
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
	for k > 1 && i.greater(k/2,k) {
		i.Exch(k/2,k)
		k = k/2
	}
}

func (i *IndexMinPQ) Insert(v int, item float64) {
	log.Println("Inserting to PQ",v)
	N := i.size() - 1
	i.QP[v] = N // last queue position for the new item
	i.PQ[N] = v
	i.Item[v] = item 
	log.Print("QP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n==================================\n")
	i.Swim(N) // swim up with the item we just added
	log.Print("AFTER SWIMMING:\nQP:",i.QP,"\nPQ:",i.PQ,"\nItem:",i.Item,"\n")

}

func (i *IndexMinPQ) DecreaseKey(v int, item float64) {
	i.Item[v] = item // add this item to the vertice v
	i.Swim(i.QP[v]) // swim the queue position up
}

func (i *IndexMinPQ) Contains(v int) bool {
	return i.QP[v] != -1
}

// resize the length of pq, qp, item arrays
func (i *IndexMinPQ) resize() {
	n := i.size()
	i.PQ = i.PQ[0:n-1]
}

func (i *IndexMinPQ) DelMin() int {
	N := i.size() - 1 
	log.Println("Deleting min N is:",N)
	min := i.PQ[0] // first value
	i.Exch(1,N) // switch the last value with the first
	i.resize() // reduce the length of the arrays and cut off the min value
	i.Sink(1) // sink the max value back into its places after selecting the first
	i.QP[i.PQ[N-1]] = -1 // shrink the QP by one
	return min
}

func (i *IndexMinPQ) size() int {
	fmt.Println("Length of PQ:",len(i.PQ))
	return len(i.QP) 
}

func (i *IndexMinPQ) IsEmpty() bool {
	return len(i.PQ) == 0
}

