package q

import (
	"fmt"
	"math/rand"
	"time"
)

var counter = 81
var a [9][9]byte
var colors = [7]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
var points = 0

// PrintMatr myComment
func PrintMatr() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", a[i][j])
		}
		fmt.Println()
	}
}

func createRand() {
	var x, y int
	x = rand.Intn(7)
	y = rand.Intn(counter)
	temp := 0
	for j := 0; j < 9; j++ {
		for k := 0; k < 9; k++ {
			if temp < y {
				if a[j][k] == '.' {
					temp++
				}
			} else if temp == y {
				if a[j][k] == '.' {
					a[j][k] = colors[x]
					//fmt.Printf("%d %d \n", j, k)
					temp++
				} else {
					continue
				}

			}
		}
	}
}

type coord struct {
	x int
	y int
}

func bfs(x1 int, y1 int, x2 int, y2 int) bool {
	var queue []coord
	var b [9][9]byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[i][j] = a[i][j]
		}
	}

	queue = append(queue, coord{x: x1, y: y1})
	// Top (just get next element, don't remove it)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		b[curr.x][curr.y] = '!'
		if curr.x == x2 && curr.y == y2 {
			return true
		}
		if curr.x > 0 && b[curr.x-1][curr.y] == '.' {
			queue = append(queue, coord{x: curr.x - 1, y: curr.y})
		}
		if curr.x < 8 && b[curr.x+1][curr.y] == '.' {
			queue = append(queue, coord{x: curr.x + 1, y: curr.y})
		}
		if curr.y > 0 && b[curr.x][curr.y-1] == '.' {
			queue = append(queue, coord{x: curr.x, y: curr.y - 1})
		}
		if curr.y < 8 && b[curr.x][curr.y+1] == '.' {
			queue = append(queue, coord{x: curr.x, y: curr.y + 1})
		}

	}
	return false
}

func addPoints(numb int) {
	newPoints := (numb - 4 + 1) * numb
	fmt.Printf("You earned %d points!\n", newPoints)
	points += newPoints
}

func checkLine(x int, y int) bool {
	var i, j, ix, iy, jx, jy int
	// check vertical
	i = x
	j = x
	for i > 0 && a[i-1][y] == a[x][y] {
		i--
	}
	for j < 8 && a[j+1][y] == a[x][y] {
		j++
	}
	if j-i+1 > 3 {
		addPoints(j - i + 1)
		for l := i; l <= j; l++ {
			a[l][y] = '.'
		}
		return true
	}

	// check horizontal
	i = y
	j = y
	for i > 0 && a[x][i-1] == a[x][y] {
		i--
	}
	for j < 8 && a[x][j+1] == a[x][y] {
		j++
	}
	if j-i+1 > 3 {
		addPoints(j - i + 1)
		for l := i; l <= j; l++ {
			a[x][l] = '.'
		}
		return true
	}

	// check diagonal 1
	ix = x
	iy = y
	jx = x
	jy = y
	for ix > 0 && iy > 0 && a[ix-1][iy-1] == a[x][y] {
		ix--
		iy--
	}
	for jx < 8 && jy < 8 && a[jx+1][jy+1] == a[x][y] {
		jx++
		jy++
	}
	if jx-ix+1 > 3 {
		addPoints(jx - ix + 1)
		for l := ix; l <= jx; l++ {
			a[l][l-x+y] = '.'
		}
		return true
	}

	// check diagonal 2
	ix = x
	iy = y
	jx = x
	jy = y
	for ix > 0 && iy < 8 && a[ix-1][iy+1] == a[x][y] {
		ix--
		iy++
	}
	for jx < 8 && jy > 0 && a[jx+1][jy-1] == a[x][y] {
		jx++
		jy--
	}
	if jx-ix+1 > 3 {
		addPoints(jx - ix + 1)
		for l := ix; l <= jx; l++ {
			a[l][x+y-l] = '.'
		}
		return true
	}
	return false
}

func main() {
	var x1, y1, x2, y2 int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			a[i][j] = '.'
		}
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 5; i++ {
		createRand()
		counter--
	}
	printMatr()

	for {
		for true {
			fmt.Printf("What ball to move? \n")
			fmt.Scanf("%d %d\n", &x1, &y1)
			if a[x1][y1] != '.' {

				break
			}
			fmt.Printf("Empty square\n")
		}

		for true {
			fmt.Printf("Where to move? \n")
			fmt.Scanf("%d %d\n", &x2, &y2)
			if a[x2][y2] != '.' {
				fmt.Printf("Ok. Will move this ball\n")
				x1 = x2
				y1 = y2
			} else {
				if bfs(x1, y1, x2, y2) {
					fmt.Printf("Great!\n")
					break
				} else {
					fmt.Printf("No path\n")
					continue
				}

			}
		}
		a[x2][y2] = a[x1][y1]
		a[x1][y1] = '.'

		if !checkLine(x2, y2) {
			for i := 0; i < 3; i++ {
				createRand()
				counter--
			}
		} else {
			fmt.Printf("Total is %d\n", points)
		}
		printMatr()
	}
}
