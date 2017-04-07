package lines

import (
	"math/rand"
	"sync"
	"time"

	"fmt"

	"github.com/MariaRogulenko/lines/api"
	"golang.org/x/net/context"
)

// Service implements the api.Game GRPC interface.
type Service struct {
	mu sync.Mutex
}

// DBComminication implements
type DBComminication struct {
	id       string
	username string
	score    int32
	table    []int32
	active   Point
}

// Point implements
type Point struct {
	x int32
	y int32
}

// NewService creates a new instance of the Service struct.
func NewService() *Service {
	return &Service{}
}

// Login implementes the api.Game GRPC interface.
func (s *Service) Login(_ context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := ReadItem(req.Id)
	if result != nil {
		return &api.LoginResponse{Id: req.Id}, nil
	}
	var arr = []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	result = &DBComminication{
		id:       req.Id,
		username: req.UserName,
		score:    0,
		table:    arr,
		active: Point{
			x: -1,
			y: -1,
		},
	}
	var counter int32 = 81
	for i := 0; i < 5; i++ {
		createRand(counter, result)
		counter--
	}
	StoreItem(result)
	return &api.LoginResponse{Id: req.Id}, nil
}

// New implementes the api.Game GRPC interface.
func (s *Service) New(_ context.Context, req *api.NewRequest) (*api.NewResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := ReadItem(req.Id)
	var arr = []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	result.table = arr
	result.active.x = -1
	result.active.y = -1
	var counter int32 = 81
	for i := 0; i < 5; i++ {
		createRand(counter, result)
		counter--
	}
	StoreItem(result)
	return &api.NewResponse{Changed: true, State: &api.State{
		Status: api.Status_READY,
		Board: &api.Board{
			CreatedBy: result.username,
			Table:     result.table,
			Active: &api.Point{
				X: result.active.x,
				Y: result.active.y,
			},
		}}}, nil
}

// GetState implementes the api.Game GRPC interface.
func (s *Service) GetState(_ context.Context, req *api.StateRequest) (*api.State, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result = ReadItem(req.Id)
	//state, ok := s.states[req.Id]
	//if !ok {
	//	return &api.State{Status: api.Status_NOT_FOUND}, nil
	//}
	return &api.State{
		Status: api.Status_READY,
		Board: &api.Board{
			CreatedBy: result.username,
			Table:     result.table,
			Active: &api.Point{
				X: result.active.x,
				Y: result.active.y,
			},
		}}, nil
}

// Move implementes the api.Game GRPC interface.
func (s *Service) Move(_ context.Context, req *api.MoveRequest) (*api.MoveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result = ReadItem(req.Id)
	//state, ok := s.states[req.Id]
	//if !ok {
	//	return &api.MoveResponse{State: &api.State{Status: api.Status_NOT_FOUND}}, nil
	//}
	var counter int32
	for i := 0; i < 81; i++ {
		if result.table[i] == 0 {
			counter++
		}
	}
	to := req.To.X*9 + req.To.Y
	if result.active.x != -1 {
		if result.table[to] > 0 {
			fmt.Println("im here 1")
			result.active.x = req.To.X
			result.active.y = req.To.Y
		} else {
			dist := Point{x: req.To.X, y: req.To.Y}
			if bfs(result, dist) {
				from := result.active.x*9 + result.active.y
				fmt.Println("im here 2")
				result.table[from], result.table[to] = result.table[to], result.table[from]
				result.active.x = -1
				result.active.y = -1
				if !checkLine(result, dist) {
					fmt.Println("im here 3")
					for i := 0; i < 3; i++ {
						createRand(counter, result)
						counter--
					}
				}
			}
		}
	} else {
		if result.table[to] > 0 {
			fmt.Println("im here 4")
			result.active = Point{x: req.To.X, y: req.To.Y}
		}
	}
	StoreItem(result)
	return &api.MoveResponse{Changed: true, State: &api.State{
		Status: api.Status_READY,
		Board: &api.Board{
			CreatedBy: result.username,
			Table:     result.table,
			Active: &api.Point{
				X: result.active.x,
				Y: result.active.y,
			},
		}}}, nil
}

func createRand(counter int32, state *DBComminication) {
	var x, y, j, k int32
	rand.Seed(time.Now().UTC().UnixNano())
	x = rand.Int31n(7)
	y = rand.Int31n(counter)
	var temp int32
	for j = 0; j < 9; j++ {
		for k = 0; k < 9; k++ {
			if temp < y {
				if state.table[j*9+k] == 0 {
					temp++
				}
			} else if temp == y {
				if state.table[j*9+k] == 0 {
					state.table[j*9+k] = x + 1
					checkLine(state, Point{x: j, y: k})
					temp++
				} else {
					continue
				}
			}
		}
	}
}

var (
	dx = []int32{-1, 1, 0, 0}
	dy = []int32{0, 0, 1, -1}
)

func bfs(state *DBComminication, to Point) bool {
	var queue []Point
	var u [9][9]bool
	from := state.active
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			u[i][j] = state.table[i*9+j] != 0
		}
	}
	queue = append(queue, from)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.x == to.x && curr.y == to.y {
			return true
		}
		for i := 0; i < 4; i++ {
			next := Point{x: curr.x + dx[i], y: curr.y + dy[i]}
			if next.x >= 0 && next.x < 9 && next.y >= 0 && next.y < 9 && !u[next.x][next.y] {
				u[next.x][next.y] = true
				queue = append(queue, next)
			}
		}
	}
	return false
}

func checkLine(state *DBComminication, curr Point) bool {
	var i, j, ix, iy, jx, jy int32
	// check vertical
	x := curr.x
	y := curr.y
	i = x
	j = x
	table := state.table
	for i > 0 && table[(i-1)*9+y] == table[x*9+y] {
		i--
	}
	for j < 8 && table[(j+1)*9+y] == table[x*9+y] {
		j++
	}
	if j-i+1 > 3 {
		for l := i; l <= j; l++ {
			table[l*9+y] = 0
		}
		return true
	}
	// check horizontal
	i = y
	j = y
	for i > 0 && table[x*9+i-1] == table[x*9+y] {
		i--
	}
	for j < 8 && table[x*9+j+1] == table[x*9+y] {
		j++
	}
	if j-i+1 > 3 {
		for l := i; l <= j; l++ {
			table[x*9+l] = 0
		}
		return true
	}
	// check diagonal 1
	ix = x
	iy = y
	jx = x
	jy = y
	for ix > 0 && iy > 0 && table[(ix-1)*9+iy-1] == table[x*9+y] {
		ix--
		iy--
	}
	for jx < 8 && jy < 8 && table[(jx+1)*9+jy+1] == table[x*9+y] {
		jx++
		jy++
	}
	if jx-ix+1 > 3 {
		for l := ix; l <= jx; l++ {
			table[l*9+l-x+y] = 0
		}
		return true
	}
	// check diagonal 2
	ix = x
	iy = y
	jx = x
	jy = y
	for ix > 0 && iy < 8 && table[(ix-1)*9+iy+1] == table[x*9+y] {
		ix--
		iy++
	}
	for jx < 8 && jy > 0 && table[(jx+1)*9+jy-1] == table[x*9+y] {
		jx++
		jy--
	}
	if jx-ix+1 > 3 {
		for l := ix; l <= jx; l++ {
			table[l*9+x+y-l] = 0
		}
		return true
	}
	return false
}
