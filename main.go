package main

import (
	"fmt"
	"strconv"
)

const INF = 1000000000

func minDistance(dist []int, visited []bool, V int) int {
	minVal := INF
	minIndex := -1
	for v := 0; v < V; v++ {
		if !visited[v] && dist[v] < minVal {
			minVal = dist[v]
			minIndex = v
		}
	}
	return minIndex
}

func dijkstra(graph [][]int, src int, V int) ([]int, []int) {
	dist := make([]int, V)
	visited := make([]bool, V)
	parent := make([]int, V)

	for i := 0; i < V; i++ {
		dist[i] = INF
		visited[i] = false
		parent[i] = -1
	}

	dist[src] = 0

	for count := 0; count < V-1; count++ {
		u := minDistance(dist, visited, V)
		if u == -1 || dist[u] == INF {
			break
		}
		visited[u] = true
		for v := 0; v < V; v++ {
			if !visited[v] && graph[u][v] != INF && dist[u] != INF {
				if dist[u]+graph[u][v] < dist[v] {
					dist[v] = dist[u] + graph[u][v]
					parent[v] = u
				}
			}
		}
	}

	return dist, parent
}

func printPath(parent []int, target int) {
	if parent[target] == -1 {
		fmt.Print(target)
		return
	}
	printPath(parent, parent[target])
	fmt.Printf(" -> %d", target)
}

func printSolution(src int, dist []int, parent []int, V int) {
	fmt.Printf("\nSource vertex: %d\n\n", src)
	fmt.Printf("%-10s%-12s%s\n", "Vertex", "Distance", "Path")
	for v := 0; v < V; v++ {
		if dist[v] == INF {
			fmt.Printf("%-10d%-12s%s\n", v, "INF", "Unreachable")
		} else {
			fmt.Printf("%-10d%-12d", v, dist[v])
			printPath(parent, v)
			fmt.Println()
		}
	}
}

func main() {
	fmt.Print("Enter number of vertices: ")
	var V int
	if _, err := fmt.Scan(&V); err != nil || V <= 0 {
		fmt.Println("Error: Invalid number of vertices")
		return
	}

	graph := make([][]int, V)
	fmt.Println("Enter adjacency matrix (use INF for no edge):")
	for i := 0; i < V; i++ {
		graph[i] = make([]int, V)
		for j := 0; j < V; j++ {
			var token string
			if _, err := fmt.Scan(&token); err != nil {
				fmt.Println("Error: Incomplete adjacency matrix")
				return
			}
			if token == "INF" || token == "inf" || token == "Infinity" {
				graph[i][j] = INF
			} else {
				weight, err := strconv.Atoi(token)
				if err != nil {
					fmt.Printf("Error: Invalid weight '%s'\n", token)
					return
				}
				if weight < 0 {
					fmt.Printf("Error: Negative edge weight detected (%d)\n", weight)
					return
				}
				graph[i][j] = weight
			}
		}
	}

	fmt.Print("Enter source vertex: ")
	var src int
	if _, err := fmt.Scan(&src); err != nil || src < 0 || src >= V {
		fmt.Println("Error: Invalid source vertex")
		return
	}

	dist, parent := dijkstra(graph, src, V)
	printSolution(src, dist, parent, V)
}
