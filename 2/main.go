package main

import (
	"fmt"
	"log"
	"math"
)

type Path struct {
	vertices []int
}



var (
	allPath []Path
)

func main() {

	var numberOfVertices int
	var numberOfEdges int
	fmt.Println("Enter number of vertices and edges: ")
	_, err := fmt.Scanf("%d %d\n", &numberOfVertices, &numberOfEdges)
	if err != nil {
		log.Fatal("Bad input...")
	}

	graph := make([][]int, numberOfVertices)
	for i := range graph {
		graph[i] = make([]int, numberOfVertices)
	}

	// Get Input
	var u, v int
	for i := 0; i < numberOfEdges; i++ {
		fmt.Println("Enter (u v)")
		_, err = fmt.Scanf("%d %d", &u, &v)
		if err != nil {
			log.Fatal("Bad input...")
		}
		// Set capacity as 1
		graph[u][v] = 1
	}

	// Define source and sink
	var source, sink int
	fmt.Println("Enter source and sink")
	_, err = fmt.Scanf("%d %d\n", &source, &sink)
	if err != nil {
		log.Fatal("Bad input...")
	}

	result := maxFlow(graph, source, sink)

	if result >=2 {
		fmt.Println("Sheep will have a nice trip! ")
		for _, path := range allPath {
			fmt.Println("Steps:" )
			for i := len(path.vertices)- 1 ; i>=0; i-- {
				fmt.Println(path.vertices[i] )
			}

		}
	}
}

func maxFlow(graph [][]int, source, sink int) int {
	numberOfVertices := len(graph)
	// Create and initialize residual graph
	residualGraph := make([][]int, numberOfVertices)
	for i := range residualGraph {
		residualGraph[i] = make([]int, numberOfVertices)
	}

	for i := range residualGraph {
		for j := range residualGraph {
			residualGraph[i][j] = graph[i][j]
		}
	}

	// Create an array for store parents of current node for determining lowest cost path
	// additionally it helps us to store path
	parent := make([]int, numberOfVertices)
	for i := range parent {
		parent[i] = -1
	}

	maxFlow := 0
	// find augment path and calculate max flow
	for existsPath(residualGraph, parent, source, sink) {
		pathFlow := math.MaxInt64
		var path Path

		// find minimum flow in path
		for v := sink ; v != source; v = parent[v] {
			u := parent[v]
			path.vertices = append(path.vertices, v)
			if pathFlow  > residualGraph[u][v] {
				pathFlow = residualGraph[u][v]
			}
		}
		path.vertices = append(path.vertices, 0)

		// Update all edges that are in the path
		for v := sink ; v != source; v = parent[v] {
			u := parent[v]
			residualGraph[u][v] = residualGraph[u][v] - pathFlow
			residualGraph[v][u] = residualGraph[v][u] + pathFlow
		}

		maxFlow += pathFlow
		allPath = append(allPath, path)
	}
	return maxFlow
}

func existsPath(residualGraph [][]int, parent []int, source int, sink int) bool {
	numberOfVertices := len(residualGraph)
	// Store visited vertices for doing bfs
	visitedVertices := make([]bool, len(residualGraph))
	for i := range visitedVertices {
		visitedVertices[i] = false
	}
	var queue []int

	queue = append(queue, source)

	visitedVertices[source] = true

	var u int

	for len(queue) != 0 {
		u = queue[0]
		queue = queue[1:]
		for v := 0; v < numberOfVertices; v++ {
			if visitedVertices[v] == false && residualGraph[u][v] > 0 {
				parent[v] = u
				visitedVertices[v] = true
				if v == sink {
					return true

				}else {
					queue = append(queue, v)
				}
			}
		}
	}
	return false
}
