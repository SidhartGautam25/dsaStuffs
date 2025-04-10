package main

import (
	"fmt"

	"github.com/SidhartGautam25/dsaStuffs/graph/dijikstra"
)

func main() {
	graph := dijikstra.Graph{
		{{To: 1, Weight: 4}, {To: 2, Weight: 1}},
		{{To: 3, Weight: 1}},
		{{To: 1, Weight: 2}, {To: 3, Weight: 5}},
		{},
	}

	start := 0
	dist := dijikstra.Dijikstra(graph, start)

	for i, d := range dist {
		fmt.Printf("Distance from %d to %d = %d\n", start, i, d)
	}
}
