package main

type Orange struct {
	value int
	x     int
	y     int
	depth int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows < 1 {
		return 0
	}
	columns := len(grid[0])
	if columns < 1 {
		return 0
	}
	orangeBasket := make([][]Orange, rows)
	goodOranges := 0
	bfsQueue := make([]Orange, rows*columns)
	for i := 0; i < rows; i++ {
		currentRow := make([]Orange, columns)
		for j := 0; j < columns; j++ {
			value := grid[i][j]
			newNode := Orange{x: i, y: j, value: value, depth: 0}
			currentRow = append(currentRow, newNode)
			if value == 2 {
				bfsQueue = append(bfsQueue, newNode)
			}
			if value == 1 {
				goodOranges++
			}
			orangeBasket[i] = currentRow
		}
	}
	if goodOranges == 0 {
		return 0
	}
	if len(bfsQueue) == 0 {
		return -1
	}
	depth := 0
	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(bfsQueue) > 0 {
		orange := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		if orange.value == 0 {
			continue
		}
		if orange.depth > depth {
			depth = orange.depth
		}
		for _, direction := range directions {
			xPrime := orange.x + direction[0]
			yPrime := orange.y + direction[1]
			if xPrime >= rows || xPrime < 0 || yPrime >= columns || yPrime < 0 {
				continue
			}
			neighbour := orangeBasket[xPrime][yPrime]
			if neighbour.value != 1 {
				continue
			}
			neighbour.depth = orange.depth + 1
			goodOranges -= 1
			neighbour.value = 2
			bfsQueue = append(bfsQueue, neighbour)
		}
	}
	if goodOranges > 0 {
		return -1
	}
	return depth
}
