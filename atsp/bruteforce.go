package atsp

import (
	"math"

	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

// pomocnicza struktur zawierająca graf
type BruteForce struct {
	graph graph.Graph
}

// funkcja zwracająca graf na którym pracuje bruteforce
func (bf *BruteForce) GetGraph() graph.Graph {
	return bf.graph
}

// funkcja tworząca nowy BruteForce
func NewBruteforce(g graph.Graph) *BruteForce {
	return &BruteForce{graph: g}
}

// funkcja rozwiązująca problem atsp dla pierwszego miasta
func (bf *BruteForce) Solve() (int, []int) {
	n := bf.graph.GetVerticesNum()
	nodes := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		nodes[i] = i + 1
	}

	bestCost := math.MaxInt32
	var bestPath []int

	var permute func([]int, int)
	permute = func(a []int, k int) {
		if k == len(a) {
			currentCost, _ := bf.graph.GetPath(0, a[0])
			for i := 0; i < len(a)-1; i++ {
				c, _ := bf.graph.GetPath(a[i], a[i+1])
				currentCost += c
			}
			co, _ := bf.graph.GetPath(a[len(a)-1], 0)
			currentCost += co

			if currentCost < bestCost {
				bestCost = currentCost
				bestPath = make([]int, len(a)+2)
				bestPath[0] = 0
				copy(bestPath[1:], a)
				bestPath[len(bestPath)-1] = 0
			}
			return
		}
		for i := k; i < len(a); i++ {
			a[k], a[i] = a[i], a[k]
			permute(a, k+1)
			a[k], a[i] = a[i], a[k]
		}
	}

	permute(nodes, 0)
	return bestCost, bestPath
}
