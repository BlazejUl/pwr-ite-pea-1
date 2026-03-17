package atsp

import (
	"fmt"

	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

type Random struct {
	graph graph.Graph
}

func (ra *Random) GetGraph() graph.Graph {
	return ra.graph
}

func NewRandom(G graph.Graph) *Random {
	return &Random{graph: G}
}

func (ra *Random) Solve(startVertex int) (int, []int) {
	visited := make([]bool, ra.graph.GetVerticesNum())
	path := make([]int, 0, ra.graph.GetVerticesNum())
	bestPath := make([]int, ra.graph.GetVerticesNum())
	bestCost := 2147483644
	//cost := 0

	visited[startVertex] = true
	path = append(path, startVertex)
	for range 10 * ra.graph.GetVerticesNum() {
		fmt.Println("l")
	}
	//cost, path = ra.BFRec(startVertex)

	return bestCost, bestPath
}
