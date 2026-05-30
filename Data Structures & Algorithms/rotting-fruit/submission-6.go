
type (
	Coords struct {
		X,Y, dist int
	}
	RottenOrange struct {
		curr *Coords
		Q []*Coords
		next, prev *RottenOrange
	}
	List struct{

	}
)

const (
	zero int = iota
	fresh
    rotten
)

func orangesRotting(grid [][]int) int {
    var maxDist = 0

	Q := []Coords{}
	var R, C = len(grid), len(grid[0])
	freshCount := 0
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++{
			if grid[i][j]==rotten {
				Q = append(Q, Coords{i,j,0})
			}else if grid[i][j]==fresh {
				freshCount++
			}
		}
	}

	for len(Q)>0{
		curr := Q[0]
		Q = Q[1:]
		maxDist = max(maxDist, curr.dist)
		freshCount = rot(grid, &curr, &Q, freshCount)
	}

	if freshCount > 0 {
		return -1
	}

	return maxDist
}

func max (a,b int)int{
	if a > b {
		return a
	}
	return b
}

func rot(grid [][]int, curr *Coords, Q *[]Coords, freshCount int) int {
	var R, C = len(grid), len(grid[0])
	var x, y, dist = curr.X, curr.Y, curr.dist

	if y+1<C && grid[x][y+1] == fresh {
		*Q = append(*Q, Coords{x, y+1, dist+1})
		grid[x][y+1]=rotten
		freshCount--
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
	//fmt.Printf("grid[x+1][y]: %d\n", grid[x+1][y])
	if x+1<R && grid[x+1][y] == fresh{
		*Q = append(*Q, Coords{x+1,y,dist+1})
		grid[x+1][y]=rotten
		freshCount--
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
	//	fmt.Printf("grid[x][y-1]: %d\n", grid[x][y-1])
	if y-1>=0 && grid[x][y-1] == fresh{
		*Q = append(*Q, Coords{x,y-1,dist+1})
		grid[x][y-1]=rotten
		freshCount--
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
//		fmt.Printf("grid[x-1][y]: %d\n", grid[x-1][y])

	if x-1>=0 && grid[x-1][y] == fresh{
		*Q = append(*Q, Coords{x-1,y,dist+1})
		grid[x-1][y]=rotten
		freshCount--
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
//	fmt.Printf("successfully added to queue: %#v\n", *Q)
	return freshCount
	//return dist + 1
}
