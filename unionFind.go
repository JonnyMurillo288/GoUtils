package utils

type Graph struct {
	edges map[int][]int
}

// default is directed, if you want an undirected addEdge(v,w) then addEdge(w,v)
func (g *Graph) AddEdge(v int, w int) {
	g.edges[v] = append(g.edges[v],w)
}

func (g *Graph) V() int {
	var len int
	for _,_ = range g.edges {
		len++
	}
	return len
}

func FindParent(p []int, v int) int {
	if p[v] == -1 {
		return v
	} else {
		return FindParent(p,p[v])
	}
}

func Union(p []int,v int, w int) []int {
	p[v] = w
	return p
}

func (g *Graph) IsCyclic() bool {
	var parent []int
	for i := 0; i < g.V(); i++ {
		parent = append(parent,-1)
	}
	for i := 0; i < g.V(); i++ {
		for _,j := range g.edges[i] {
			x := FindParent(parent,i)
			y := FindParent(parent,j)
			if x == y {
				return true
			}
			Union(parent,x,y)
		}
	}
	return false
}