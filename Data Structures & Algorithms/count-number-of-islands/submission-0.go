func numIslands(grid [][]byte) int {
    totalIslands := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j]=='1'{
                totalIslands++
                discoverIsland(grid, i, j)
            }
        }
    }
    return totalIslands
}

func discoverIsland(grid [][]byte, i, j int) {
    grid[i][j]='0'

    if i-1 >= 0 && grid[i-1][j]=='1' {
        discoverIsland(grid, i-1, j)
    }

    if j+1 < len(grid[0]) && grid[i][j+1]=='1' {
        discoverIsland(grid, i, j+1)
    }

    if i+1 < len(grid) && grid[i+1][j]=='1' {
        discoverIsland(grid, i+1, j)
    }

    if j-1 >= 0 && grid[i][j-1]=='1' {
        discoverIsland(grid, i, j-1)
    }
}
