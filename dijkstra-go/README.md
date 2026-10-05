# Dijkstra's Algorithm in Go

## Algorithm Name
Dijkstra's Shortest Path Algorithm

## Short Description
Dijkstra's algorithm finds the shortest path from a single source vertex to all other vertices in a weighted graph with non-negative edge weights.

### Core Steps:
1. **Initialize**: Set `dist[src] = 0` and all other `dist[v] = INF`. Mark all vertices as unvisited (`visited[v] = false`) and set `parent[v] = -1`.
2. **Select Minimum**: Pick the unvisited vertex `u` with the smallest tentative distance.
3. **Mark Visited**: Mark `u` as visited.
4. **Relax**: For each unvisited adjacent vertex `v`, if `dist[u] + weight(u, v) < dist[v]`, update `dist[v] = dist[u] + weight(u, v)` and set `parent[v] = u`.
5. **Repeat**: Repeat until all reachable vertices are processed.

---

## Complexity
- **Time Complexity**: $\mathcal{O}(V^2)$ — Selecting the minimum-distance unvisited vertex using a linear search takes $\mathcal{O}(V)$, and relaxing its neighbors using the adjacency matrix row takes $\mathcal{O}(V)$, repeated for all $V$ vertices.
- **Space Complexity**: $\mathcal{O}(V^2)$ — Storing the graph using a $V \times V$ adjacency matrix.

---

## How to Run

```bash
# Using input redirection
go run main.go < sample_input.txt

# Or interactively
go run main.go
```

On Windows PowerShell:
```powershell
Get-Content sample_input.txt | go run main.go
```

---

## Input Format
1. Number of vertices ($V$)
2. $V \times V$ weighted adjacency matrix (enter integer weights or `INF` for absence of edge)
3. Source vertex ($0$ to $V-1$)

---

## Output Format
Displays the source vertex followed by a table with:
- `Vertex`: Target vertex index
- `Distance`: Shortest distance from source (or `INF` if unreachable)
- `Path`: Reconstructed shortest path (or `Unreachable`)

---

## Sample Input (`sample_input.txt`)
```text
5
0 10 3 INF INF
INF 0 1 2 INF
INF 4 0 8 2
INF INF INF 0 7
INF INF INF 9 0
0
```

---

## Sample Output
```text
Source vertex: 0

Vertex    Distance    Path
0         0           0
1         7           0 -> 2 -> 1
2         3           0 -> 2
3         9           0 -> 2 -> 1 -> 3
4         5           0 -> 2 -> 4
```
