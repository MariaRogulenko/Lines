package lines

import (
	"fmt"
	"sync"

	"github.com/MariaRogulenko/lines/api"
	"golang.org/x/net/context"
)

// Service implementes the api.Game GRPC interface.
type Service struct {
	mu     sync.Mutex
	maxID  int
	states map[string]*api.State
}

// NewService creates a new instance of the Service struct.
func NewService() *Service {
	return &Service{states: map[string]*api.State{}}
}

// New implementes the api.Game GRPC interface.
func (s *Service) New(_ context.Context, req *api.NewRequest) (*api.NewResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxID++
	id := fmt.Sprintf("%d", s.maxID)
	s.states[id] = &api.State{
		Status: api.Status_READY,
		Board: &api.Board{
			CreatedBy: req.UserName,
			Table: []int32{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
	}
	return &api.NewResponse{Id: id}, nil
}

// GetState implementes the api.Game GRPC interface.
func (s *Service) GetState(_ context.Context, req *api.StateRequest) (*api.State, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	state, ok := s.states[req.Id]
	if !ok {
		return &api.State{Status: api.Status_NOT_FOUND}, nil
	}
	return state, nil
}

// Move implementes the api.Game GRPC interface.
func (s *Service) Move(_ context.Context, req *api.MoveRequest) (*api.MoveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	state, ok := s.states[req.Id]
	if !ok {
		return &api.MoveResponse{State: &api.State{Status: api.Status_NOT_FOUND}}, nil
	}
	from := req.From.X*9 + req.From.Y
	to := req.To.X*9 + req.To.Y
	state.Board.Table[from], state.Board.Table[to] = state.Board.Table[to], state.Board.Table[from]
	return &api.MoveResponse{Changed: true, State: state}, nil
}
