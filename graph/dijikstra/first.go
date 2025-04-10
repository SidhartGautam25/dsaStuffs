package dijikstra

import (
	"container/heap"
	"math"
)

/*

1. Starting me saare hi node infinity distance pe hai ye maanke chalenge
2. Src node ko distance 0 pe maanenge
3. ab us src node ke jitne children hai unke distance ko update krenge
       --> distance[children]=distance[currNode]+weight
	   agr ye value kam hai us distance se jo abhi hai uske or src ke bich
	   to usko update kr denge

	   Aisa hi hr children ke sath krenge
4. iske baad ab us node pe jayenge jo sbse najdik hai src se abhi tk ke calculation
   ke hisab se
          --> Is chiz ke lie hmlog priority queue ka use kr sakte hai

5. Aisa krte rhenge jbtk hr node ke lie ye calculation na ho jaye


*/

type Edge struct {
	To     int
	Weight int
}

type Graph [][]Edge

// Item is used in the priority queue
type Item struct {
	node     int
	distance int
	index    int
}

// Start of Priority Queue Implementation
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
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

// End of implementation

func Dijikstra(graph Graph, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[start] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		u := curr.node

		for _, edge := range graph[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{node: v, distance: dist[v]})
			}

		}
	}

	return dist
}
