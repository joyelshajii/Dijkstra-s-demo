# Dijkstra's Shortest Path Algorithm (Go)

## Algorithm Name
**Dijkstra's Shortest Path Algorithm** (Single-Source Shortest Path)

---

## Short Description
Dijkstra's Algorithm is a greedy graph search algorithm that finds the shortest path from a single source vertex to all other vertices in a weighted graph with **non-negative edge weights**.

### How the Algorithm Works:
1. **Initialization**:
   - Set the distance to the source vertex as `0` (`dist[src] = 0`).
   - Set the tentative distance to all other vertices as infinity (`dist[v] = INF`).
   - Initialize a `visited` set/array to track vertices whose shortest distance is finalized (`visited[v] = false`).
   - Initialize a `parent` array to track the predecessor of each vertex for path reconstruction (`parent[v] = -1`).

2. **Greedy Vertex Selection**:
   - Among all vertices that have not yet been visited, select the vertex $u$ with the minimum tentative distance:
     $$\min \{ \text{dist}[u] \mid \text{visited}[u] = \text{false} \}$$
   - If the minimum distance vertex has tentative distance equal to `INF`, all remaining unvisited vertices are unreachable, and the algorithm terminates.

3. **Mark as Visited**:
   - Mark vertex $u$ as visited (`visited[u] = true`). Its shortest distance from the source is now finalized.

4. **Edge Relaxation**:
   - For every adjacent vertex $v$ of $u$:
     - Check if $v$ is unvisited and an edge $(u, v)$ exists.
     - If the path through $u$ is shorter than the currently recorded distance to $v$:
       $$\text{dist}[u] + \text{weight}(u, v) < \text{dist}[v]$$
     - Update $\text{dist}[v] = \text{dist}[u] + \text{weight}(u, v)$ and update $\text{parent}[v] = u$.

5. **Repeat**:
   - Repeat steps 2 to 4 until all reachable vertices are processed.

6. **Path Reconstruction**:
   - Backtrack from each destination vertex $v$ to the source vertex using the `parent` array to reconstruct the full path sequence.

---

## Complexity Analysis

- **Time Complexity: $O(V^2)$**
  - Finding the unvisited vertex with the minimum tentative distance requires a linear scan over all $V$ vertices, taking $O(V)$ time per vertex.
  - Repeating this for all $V$ vertices takes $V \times O(V) = O(V^2)$ time.
  - Examining and relaxing all neighbors of vertex $u$ across its row in the adjacency matrix takes $O(V)$ time per vertex, totaling $V \times O(V) = O(V^2)$ time.
  - Therefore, the total time complexity is:
    $$O(V^2) + O(V^2) = O(V^2)$$

- **Space Complexity: $O(V^2)$**
  - Storing the graph as a $V \times V$ adjacency matrix requires $O(V^2)$ space.
  - The auxiliary arrays (`dist`, `visited`, and `parent`) each require $O(V)$ space.
  - Total space complexity:
    $$O(V^2) + O(V) = O(V^2)$$

---

## Project Structure

```text
dijkstra-go/
├── main.go            # Dijkstra's algorithm implementation and CLI interface
├── README.md          # Algorithm documentation, complexity, and instructions
└── sample_input.txt   # Sample input graph with 5 vertices
```

---

## How to Run the Program

Prerequisites: Go (version 1.18 or newer recommended).

### Option 1: Run with an input file as argument
```bash
go run main.go sample_input.txt
```

### Option 2: Run using standard input redirection
```bash
go run main.go < sample_input.txt
```

### Option 3: Run interactively
```bash
go run main.go
```
The program will prompt you to enter the number of vertices, the adjacency matrix row by row (using `INF` for absent edges), and the source vertex.

### Option 4: Compile to a binary executable
```bash
go build -o dijkstra main.go
./dijkstra sample_input.txt
```
*(On Windows: `.\dijkstra.exe sample_input.txt`)*

---

## Input Format

The program expects three inputs:
1. **Number of vertices ($V$)**: An integer greater than 0. Vertices are indexed from `0` to `V - 1`.
2. **Adjacency Matrix ($V \times V$)**:
   - Non-negative integer weights for valid edges.
   - `0` along the diagonal (distance from a vertex to itself).
   - `INF` (or `inf`) representing the absence of an edge between vertices.
3. **Source vertex**: An integer between `0` and `V - 1`.

---

## Output Format

The output displays:
1. The chosen **Source vertex**.
2. A table with three columns:
   - `Vertex`: The target vertex index.
   - `Distance`: The shortest path distance from the source (or `INF` if unreachable).
   - `Path`: The reconstructed shortest path sequence (e.g. `0 -> 2 -> 1`), or `Unreachable` if no path exists.

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

---

## Step-by-Step Algorithm Trace (Sample Graph)

- **Source Vertex**: `0`
- **Initial State**:
  - `dist = [0, INF, INF, INF, INF]`
  - `visited = [F, F, F, F, F]`
  - `parent = [-1, -1, -1, -1, -1]`

| Iteration | Min Unvisited ($u$) | Action & Edge Relaxation | Updated `dist` Array | Updated `parent` Array |
| :---: | :---: | :--- | :--- | :--- |
| **Start** | - | Initialize source `dist[0] = 0` | `[0, INF, INF, INF, INF]` | `[-1, -1, -1, -1, -1]` |
| **1** | **0** (`dist=0`) | Relax edges $(0,1): 10$, $(0,2): 3$ | `[0, 10, 3, INF, INF]` | `[-1, 0, 0, -1, -1]` |
| **2** | **2** (`dist=3`) | Relax edges $(2,1): 3+4=7 < 10$, $(2,3): 3+8=11$, $(2,4): 3+2=5$ | `[0, 7, 3, 11, 5]` | `[-1, 2, 0, 2, 2]` |
| **3** | **4** (`dist=5`) | Relax edge $(4,3): 5+9=14 > 11$ (no change) | `[0, 7, 3, 11, 5]` | `[-1, 2, 0, 2, 2]` |
| **4** | **1** (`dist=7`) | Relax edge $(1,3): 7+2=9 < 11$ | `[0, 7, 3, 9, 5]` | `[-1, 2, 0, 1, 2]` |
| **5** | **3** (`dist=9`) | Relax edge $(3,4): 9+7=16 > 5$ (no change) | `[0, 7, 3, 9, 5]` | `[-1, 2, 0, 1, 2]` |

### Path Reconstruction:
- Vertex 0: Source $\rightarrow$ `0` (Cost: 0)
- Vertex 1: $1 \leftarrow 2 \leftarrow 0 \rightarrow$ `0 -> 2 -> 1` (Cost: 7)
- Vertex 2: $2 \leftarrow 0 \rightarrow$ `0 -> 2` (Cost: 3)
- Vertex 3: $3 \leftarrow 1 \leftarrow 2 \leftarrow 0 \rightarrow$ `0 -> 2 -> 1 -> 3` (Cost: 9)
- Vertex 4: $4 \leftarrow 2 \leftarrow 0 \rightarrow$ `0 -> 2 -> 4` (Cost: 5)
