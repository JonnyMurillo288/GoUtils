package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Jonnymurillo288/go-runner/edgeWeightedGraphs/graphs"
)



func TestIndexMinPQ(t *testing.T) {
	g := graphs.NewWeightDigraph(0)

	file, err := os.Open("./testDWG.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var ind int
	for scanner.Scan() {
		ind++
		line := strings.Split(scanner.Text()," ")
		p1, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		p2, _ := strconv.Atoi(line[1])
		w, _ := strconv.ParseFloat(line[2],64)
		e := graphs.Edge{
			V: p1,
			W: p2,
			Weight: w,
		}
		g.AddEdge(e)
	}
	fmt.Println(g.Adj)
	d := graphs.NewDijkstras(g,0)
	for i,edge := range d.EdgeTo {
		fmt.Printf("%v --> %v : %v\n",i,edge,d.DistTo[i])
	}

}