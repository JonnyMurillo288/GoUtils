package utils

type Graph struct {
	edges map[int][]int
}

// default is directed, if you want an undirected addEdge(v,w) then addEdge(w,v)
func (g *Graph) addEdge(v int, w int) {
	g.edges[v] = append(g.edges[v],w)
}

func (g *Graph) V() int {
	var len int
	for _,_ = range g.edges {
		len++
	}
	return len
}

func findParent(p []int, v int) int {
	if p[v] == -1 {
		return v
	} else {
		return findParent(p,p[v])
	}
}

func union(p []int,v int, w int) []int {
	p[v] = w
	return p
}

func (g *Graph) isCyclic() bool {
	var parent []int
	for i := 0; i < g.V(); i++ {
		parent = append(parent,-1)
	}
	for i := 0; i < g.V(); i++ {
		for _,j := range g.edges[i] {
			x := findParent(parent,i)
			y := findParent(parent,j)
			if x == y {
				return true
			}
			union(parent,x,y)
		}
	}
	return false
}