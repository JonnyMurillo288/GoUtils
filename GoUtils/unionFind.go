package utils

type Graph struct {
	Edges map[int][]int
	Parent []int
}

// default is directed, if you want an undirected addEdge(v,w) then addEdge(w,v)
func (g *Graph) AddEdge(v int, w int) {
	g.Edges[v] = append(g.Edges[v],w)
}

func (g *Graph) V() int {
	var len int
	for _,_ = range g.Edges {
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

func (g *Graph) connected(v int, w int) bool {
	// add to graph then check union
	g.AddEdge(v,w)
	return g.IsCyclic()
}

func (g *Graph) Union(v int, w int) {
	g.Parent[v] = w
}

func (g *Graph) IsCyclic() bool {
	g.Parent = []int{}
	for i := 0; i < g.V(); i++ {
		g.Parent = append(g.Parent,-1)
	}
	for i := 0; i < g.V(); i++ {
		for _,j := range g.Edges[i] {
			x := FindParent(g.Parent,i)
			y := FindParent(g.Parent,j)
			if x == y {
				return true
			}
			g.Union(x,y)
		}
	}
	return false
}