package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// INF represents infinity (absence of an edge or an unreachable vertex).
// 10^9 is chosen as a sufficiently large value that prevents integer overflow
// during distance additions in the relaxation step.
const INF = 1000000000

// minDistance finds the unvisited vertex with the minimum tentative distance.
// This implements the greedy selection step of Dijkstra's algorithm.
//
// Time Complexity: O(V) via linear scan of unvisited vertices.
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

// dijkstra computes the shortest paths from the source vertex to all vertices
// in a weighted directed or undirected graph represented by an adjacency matrix.
//
// Parameters:
//   - graph: V x V adjacency matrix where graph[u][v] is edge weight (or INF)
//   - src:   index of the source vertex (0 to V-1)
//   - V:     number of vertices
//
// Returns:
//   - dist:   slice of size V with shortest distance from src to each vertex
//   - parent: slice of size V storing previous vertices for path reconstruction
//
// Complexity:
//   - Time Complexity:  O(V^2)
//   - Space Complexity: O(V^2) for the adjacency matrix, O(V) auxiliary space
func dijkstra(graph [][]int, src int, V int) ([]int, []int) {
	dist := make([]int, V)
	visited := make([]bool, V)
	parent := make([]int, V)

	// Step 1: Initialize all distances as INF, visited as false, and parent as -1
	for i := 0; i < V; i++ {
		dist[i] = INF
		visited[i] = false
		parent[i] = -1
	}

	// Distance from the source to itself is 0
	dist[src] = 0

	// Step 2: Iterate V times to process all vertices
	for count := 0; count < V; count++ {
		// Find the unvisited vertex with the minimum tentative distance
		u := minDistance(dist, visited, V)

		// If no unvisited vertex is reachable, stop early (handles disconnected graphs)
		if u == -1 || dist[u] == INF {
			break
		}

		// Mark vertex u as visited (its shortest distance is now finalized)
		visited[u] = true

		// Relax all adjacent vertices of vertex u
		for v := 0; v < V; v++ {
			// Relaxation conditions:
			// 1. Vertex v is not yet visited
			// 2. An edge exists from u to v (graph[u][v] != INF)
			// 3. Distance to u is finite (dist[u] != INF, avoids overflow)
			// 4. Path through u is strictly shorter than current dist[v]
			if !visited[v] && graph[u][v] != INF && dist[u] != INF {
				if dist[u]+graph[u][v] < dist[v] {
					dist[v] = dist[u] + graph[u][v]
					parent[v] = u // Record predecessor for shortest path reconstruction
				}
			}
		}
	}

	return dist, parent
}

// getPath reconstructs the sequence of vertices from the source to the target
// vertex by backtracking using the parent slice.
func getPath(parent []int, target int) []int {
	var path []int
	curr := target

	for curr != -1 {
		path = append(path, curr)
		curr = parent[curr]
	}

	// Reverse the path to get [source, ..., target] order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// formatPath returns the shortest path as a string "0 -> 2 -> 1",
// or "Unreachable" if no path exists.
func formatPath(parent []int, dist []int, target int) string {
	if dist[target] == INF {
		return "Unreachable"
	}

	path := getPath(parent, target)
	var parts []string
	for _, v := range path {
		parts = append(parts, strconv.Itoa(v))
	}
	return strings.Join(parts, " -> ")
}

// printPath prints the shortest path from the source to a specific target vertex.
func printPath(parent []int, dist []int, target int) {
	fmt.Println(formatPath(parent, dist, target))
}

// printSolution displays the final shortest distances and paths in tabular format.
func printSolution(src int, dist []int, parent []int, V int) {
	fmt.Printf("Source vertex: %d\n\n", src)
	fmt.Printf("%-10s%-12s%s\n", "Vertex", "Distance", "Path")

	for v := 0; v < V; v++ {
		if dist[v] == INF {
			fmt.Printf("%-10d%-12s%s\n", v, "INF", "Unreachable")
		} else {
			pathStr := formatPath(parent, dist, v)
			fmt.Printf("%-10d%-12d%s\n", v, dist[v], pathStr)
		}
	}
}

// isInteractiveTerminal checks whether stdin is connected to an interactive terminal.
func isInteractiveTerminal() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func main() {
	var reader io.Reader = os.Stdin
	interactive := false

	// Optional: Accept input file as command-line argument (e.g., go run main.go sample_input.txt)
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Unable to open file '%s': %v\n", os.Args[1], err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else if isInteractiveTerminal() {
		interactive = true
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// 1. Read number of vertices (V)
	if interactive {
		fmt.Print("Enter number of vertices: ")
	}
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Error: Missing number of vertices in input.")
		os.Exit(1)
	}
	V, err := strconv.Atoi(scanner.Text())
	if err != nil || V <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Invalid number of vertices '%s'. Must be a positive integer.\n", scanner.Text())
		os.Exit(1)
	}

	// 2. Read V x V adjacency matrix
	graph := make([][]int, V)
	for i := 0; i < V; i++ {
		graph[i] = make([]int, V)
	}

	if interactive {
		fmt.Println("Enter adjacency matrix (enter integer weights or 'INF' for no edge):")
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if !scanner.Scan() {
				fmt.Fprintf(os.Stderr, "Error: Incomplete matrix. Expected %d entries for a %dx%d matrix, but reached end of input at row %d, col %d.\n", V*V, V, V, i, j)
				os.Exit(1)
			}
			token := scanner.Text()
			upperToken := strings.ToUpper(token)

			if upperToken == "INF" || upperToken == "INFINITY" {
				graph[i][j] = INF
			} else {
				weight, err := strconv.Atoi(token)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: Invalid weight token '%s' at row %d, col %d. Expected an integer or 'INF'.\n", token, i, j)
					os.Exit(1)
				}
				// Dijkstra's algorithm assumes non-negative edge weights
				if weight < 0 {
					fmt.Fprintf(os.Stderr, "Error: Negative edge weight (%d) detected at row %d, col %d. Dijkstra's algorithm does not support negative weights.\n", weight, i, j)
					os.Exit(1)
				}
				graph[i][j] = weight
			}
		}
	}

	// 3. Read source vertex (src)
	if interactive {
		fmt.Printf("Enter source vertex (0 to %d): ", V-1)
	}
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Error: Missing source vertex in input.")
		os.Exit(1)
	}
	src, err := strconv.Atoi(scanner.Text())
	if err != nil || src < 0 || src >= V {
		fmt.Fprintf(os.Stderr, "Error: Invalid source vertex '%s'. Must be an integer between 0 and %d.\n", scanner.Text(), V-1)
		os.Exit(1)
	}

	// Execute Dijkstra's Algorithm
	dist, parent := dijkstra(graph, src, V)

	// Output results
	printSolution(src, dist, parent, V)
}
