package main

import (
	"container/heap"
	"fmt"
)

type State struct {
	x, y int
	step int
	keys uint
}

/////////////////////////////////// PriorityQueue Start
// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest
	return pq[i].step < pq[j].step
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	state := x.(*State)
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	*pq = old[0 : n-1]
	return state
}

////////////////////////////////// PriorityQueue End

func haveKey(keys uint, lock uint) bool {
	return keys&(1<<lock) != 0
}

func isValid(x int, y int, keys uint, grid []string) bool {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != '#' {
		if grid[x][y] >= 'A' && grid[x][y] <= 'F' {
			return haveKey(keys, uint(grid[x][y]-'A'))
		}
		return true
	}
	return false
}

func bfs(startX int, startY int, grid []string, sizeOfKey uint) int {

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	pq := PriorityQueue{}
	var vis [100][100][1 << 6]bool
	state := &State{step: 0, x: startX, y: startY, keys: 0}
	heap.Init(&pq)
	heap.Push(&pq, state)
	vis[startX][startY][0] = true
	for {
		if len(pq) == 0 {
			break
		}
		state := heap.Pop(&pq).(*State)
		if state.keys == ((1 << sizeOfKey) - 1) {
			return state.step
		}
		for i := 0; i < 4; i++ {
			newX := state.x + dx[i]
			newY := state.y + dy[i]
			if isValid(newX, newY, state.keys, grid) && !vis[newX][newY][state.keys] {
				newKeys := state.keys
				c := grid[newX][newY]
				if c >= 'a' && c <= 'f' {
					newKeys |= (1 << (c - 'a'))
				}
				newState := &State{step: state.step + 1, x: newX, y: newY, keys: newKeys}
				pq.Push(newState)
				vis[newX][newY][newKeys] = true
			}
		}
	}
	return -1
}

func shortestPathAllKeys(grid []string) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	startX, startY := -1, -1
	var sizeOfKey uint = 0

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '@' {
				startX = i
				startY = j
			} else if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				sizeOfKey++
			}
		}
	}
	return bfs(startX, startY, grid, sizeOfKey)
}

func main() {
	grid := []string{"@.a.#", "###.#", "b.A.B"}
	fmt.Println(shortestPathAllKeys(grid))
	grid2 := []string{"@..aA", "..B#.", "....b"}
	fmt.Println(shortestPathAllKeys(grid2))
	grid3 := []string{"@...a", ".###A", "b.BCc"}
	fmt.Println(shortestPathAllKeys(grid3))
	grid4 := []string{"@abcdeABCDEFf"}
	fmt.Println(shortestPathAllKeys(grid4))
}
