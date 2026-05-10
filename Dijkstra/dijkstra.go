package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Graph map[string]map[string]int

type Item struct {
	node     string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = math.MaxInt
	}
	distances[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		currentNode := current.node
		currentDist := current.priority

		for neighbor, weight := range graph[currentNode] {
			distance := currentDist + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(&pq, &Item{node: neighbor, priority: distance})
			}
		}
	}

	return distances
}

func main() {
	graph := Graph{
		"A": {"B": 4, "C": 2},
		"B": {"A": 4, "C": 1, "D": 5},
		"C": {"A": 2, "B": 1, "D": 8, "E": 10},
		"D": {"B": 5, "C": 8, "E": 2, "Z": 6},
		"E": {"C": 10, "D": 2, "Z": 3},
		"Z": {"D": 6, "E": 3},
	}

	start := "A"
	distances := Dijkstra(graph, start)

	fmt.Printf("Menores distâncias a partir do nó %s:\n", start)
	for node, distance := range distances {
		fmt.Printf("%s: %d\n", node, distance)
	}
}
