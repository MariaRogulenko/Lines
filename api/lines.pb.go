// Code generated by protoc-gen-go.
// source: api/lines.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/lines.proto

It has these top-level messages:
	NewRequest
	NewResponse
	StateRequest
	State
	MoveRequest
	MoveResponse
	Board
	Point
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_NOT_FOUND Status = 0
	Status_READY     Status = 1
	Status_GAME_OVER Status = 2
)

var Status_name = map[int32]string{
	0: "NOT_FOUND",
	1: "READY",
	2: "GAME_OVER",
}
var Status_value = map[string]int32{
	"NOT_FOUND": 0,
	"READY":     1,
	"GAME_OVER": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NewRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *NewRequest) Reset()                    { *m = NewRequest{} }
func (m *NewRequest) String() string            { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()               {}
func (*NewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *NewRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NewResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *NewResponse) Reset()                    { *m = NewResponse{} }
func (m *NewResponse) String() string            { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()               {}
func (*NewResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StateRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StateRequest) Reset()                    { *m = StateRequest{} }
func (m *StateRequest) String() string            { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()               {}
func (*StateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type State struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=lines.api.Status" json:"status,omitempty"`
	Board  *Board `protobuf:"bytes,2,opt,name=board" json:"board,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *State) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NOT_FOUND
}

func (m *State) GetBoard() *Board {
	if m != nil {
		return m.Board
	}
	return nil
}

type MoveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	To *Point `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
}

func (m *MoveRequest) Reset()                    { *m = MoveRequest{} }
func (m *MoveRequest) String() string            { return proto.CompactTextString(m) }
func (*MoveRequest) ProtoMessage()               {}
func (*MoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MoveRequest) GetTo() *Point {
	if m != nil {
		return m.To
	}
	return nil
}

type MoveResponse struct {
	Changed bool   `protobuf:"varint,1,opt,name=changed" json:"changed,omitempty"`
	State   *State `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
}

func (m *MoveResponse) Reset()                    { *m = MoveResponse{} }
func (m *MoveResponse) String() string            { return proto.CompactTextString(m) }
func (*MoveResponse) ProtoMessage()               {}
func (*MoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MoveResponse) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
}

func (m *MoveResponse) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type Board struct {
	CreatedBy  string  `protobuf:"bytes,1,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	Score      int32   `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	Table      []int32 `protobuf:"varint,3,rep,packed,name=table" json:"table,omitempty"`
	Active     *Point  `protobuf:"bytes,4,opt,name=active" json:"active,omitempty"`
	NextColors []int32 `protobuf:"varint,5,rep,packed,name=next_colors,json=nextColors" json:"next_colors,omitempty"`
}

func (m *Board) Reset()                    { *m = Board{} }
func (m *Board) String() string            { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()               {}
func (*Board) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Board) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Board) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Board) GetTable() []int32 {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *Board) GetActive() *Point {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Board) GetNextColors() []int32 {
	if m != nil {
		return m.NextColors
	}
	return nil
}

type Point struct {
	X int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Point) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*NewRequest)(nil), "lines.api.NewRequest")
	proto.RegisterType((*NewResponse)(nil), "lines.api.NewResponse")
	proto.RegisterType((*StateRequest)(nil), "lines.api.StateRequest")
	proto.RegisterType((*State)(nil), "lines.api.State")
	proto.RegisterType((*MoveRequest)(nil), "lines.api.MoveRequest")
	proto.RegisterType((*MoveResponse)(nil), "lines.api.MoveResponse")
	proto.RegisterType((*Board)(nil), "lines.api.Board")
	proto.RegisterType((*Point)(nil), "lines.api.Point")
	proto.RegisterEnum("lines.api.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Game service

type GameClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error)
	GetState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*State, error)
	Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error) {
	out := new(NewResponse)
	err := grpc.Invoke(ctx, "/lines.api.Game/New", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := grpc.Invoke(ctx, "/lines.api.Game/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := grpc.Invoke(ctx, "/lines.api.Game/Move", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Game service

type GameServer interface {
	New(context.Context, *NewRequest) (*NewResponse, error)
	GetState(context.Context, *StateRequest) (*State, error)
	Move(context.Context, *MoveRequest) (*MoveResponse, error)
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.api.Game/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.api.Game/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetState(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.api.Game/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Move(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lines.api.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Game_New_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Game_GetState_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Game_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/lines.proto",
}

func init() { proto.RegisterFile("api/lines.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6e, 0x12, 0x41,
	0x14, 0xc6, 0x9d, 0x85, 0x45, 0xf6, 0x80, 0x95, 0x4e, 0x2a, 0xdd, 0x60, 0xab, 0x64, 0x4c, 0x4c,
	0xe5, 0x02, 0x52, 0xbc, 0xb2, 0x37, 0xa6, 0x58, 0xe4, 0xc6, 0x42, 0x9d, 0x5a, 0x13, 0x7b, 0x43,
	0x06, 0x98, 0xe0, 0x24, 0x30, 0x83, 0xbb, 0xc3, 0xbf, 0x18, 0x6f, 0x7c, 0x05, 0xaf, 0x7d, 0x2a,
	0x5f, 0xc1, 0xe7, 0x30, 0x66, 0x66, 0x76, 0x2b, 0x29, 0xf6, 0xf2, 0x7c, 0xdf, 0x99, 0x6f, 0x7f,
	0x33, 0xe7, 0x2c, 0x3c, 0x64, 0x33, 0xd1, 0x98, 0x08, 0xc9, 0xe3, 0xfa, 0x2c, 0x52, 0x5a, 0xe1,
	0xc0, 0x15, 0x6c, 0x26, 0x2a, 0x07, 0x63, 0xa5, 0xc6, 0x13, 0xde, 0x30, 0x2d, 0x4c, 0x4a, 0xa5,
	0x99, 0x16, 0x4a, 0x26, 0x8d, 0xe4, 0x15, 0x40, 0x97, 0x2f, 0x29, 0xff, 0x32, 0xe7, 0xb1, 0xc6,
	0x8f, 0x21, 0x98, 0xc7, 0x3c, 0xea, 0x4b, 0x36, 0xe5, 0x21, 0xaa, 0xa2, 0xa3, 0x80, 0xe6, 0x8d,
	0xd0, 0x65, 0x53, 0x8e, 0x77, 0xc0, 0x13, 0xa3, 0xd0, 0xb3, 0xaa, 0x27, 0x46, 0xe4, 0x10, 0x0a,
	0xf6, 0x68, 0x3c, 0x53, 0x32, 0x4e, 0x6d, 0x74, 0x63, 0x3f, 0x81, 0xe2, 0xa5, 0x66, 0x9a, 0xa7,
	0xd9, 0xb7, 0xfd, 0x6b, 0xf0, 0xad, 0x8f, 0x5f, 0x40, 0x2e, 0xd6, 0x4c, 0xcf, 0x63, 0x6b, 0xee,
	0x34, 0x77, 0xeb, 0x37, 0xf0, 0xf5, 0x4b, 0x6b, 0xd0, 0xa4, 0x01, 0x3f, 0x07, 0x7f, 0xa0, 0x58,
	0xe4, 0x28, 0x0a, 0xcd, 0xd2, 0x46, 0x67, 0xcb, 0xe8, 0xd4, 0xd9, 0xe4, 0x35, 0x14, 0xce, 0xd5,
	0xe2, 0xae, 0x4f, 0xe3, 0x2a, 0x78, 0x5a, 0x85, 0x99, 0xad, 0x8c, 0x0b, 0x25, 0xa4, 0xa6, 0x9e,
	0x56, 0xe4, 0x02, 0x8a, 0x2e, 0x20, 0xb9, 0x5c, 0x08, 0xf7, 0x87, 0x9f, 0x99, 0x1c, 0x73, 0x17,
	0x93, 0xa7, 0x69, 0x69, 0x90, 0x0c, 0x1c, 0xff, 0x0f, 0x92, 0xbb, 0xbe, 0xb3, 0xc9, 0x4f, 0x04,
	0xbe, 0x65, 0xc4, 0x87, 0x00, 0xc3, 0x88, 0x33, 0xcd, 0x47, 0xfd, 0xc1, 0x3a, 0xa1, 0x0a, 0x12,
	0xa5, 0xb5, 0xc6, 0x7b, 0xe0, 0xc7, 0x43, 0x15, 0xb9, 0x40, 0x9f, 0xba, 0xc2, 0xa8, 0x9a, 0x0d,
	0x26, 0x3c, 0xcc, 0x54, 0x33, 0x46, 0xb5, 0x05, 0x3e, 0x82, 0x1c, 0x1b, 0x6a, 0xb1, 0xe0, 0x61,
	0xf6, 0x8e, 0xcb, 0x24, 0x3e, 0x7e, 0x0a, 0x05, 0xc9, 0x57, 0xba, 0x3f, 0x54, 0x13, 0x15, 0xc5,
	0xa1, 0x6f, 0x53, 0xc0, 0x48, 0x6f, 0xac, 0x42, 0x9e, 0x81, 0x6f, 0x4f, 0xe0, 0x22, 0xa0, 0x95,
	0xa5, 0xf2, 0x29, 0x5a, 0x99, 0x6a, 0x9d, 0x90, 0xa0, 0x75, 0xed, 0x18, 0x72, 0x6e, 0x22, 0xf8,
	0x01, 0x04, 0xdd, 0xde, 0x87, 0xfe, 0xdb, 0xde, 0x55, 0xf7, 0xac, 0x74, 0x0f, 0x07, 0xe0, 0xd3,
	0xf6, 0xe9, 0xd9, 0xa7, 0x12, 0x32, 0x4e, 0xe7, 0xf4, 0xbc, 0xdd, 0xef, 0x7d, 0x6c, 0xd3, 0x92,
	0xd7, 0xfc, 0x83, 0x20, 0xdb, 0x31, 0xeb, 0xf3, 0x0e, 0x32, 0x5d, 0xbe, 0xc4, 0x8f, 0x36, 0x10,
	0xff, 0x6d, 0x5e, 0xa5, 0x7c, 0x5b, 0x76, 0x0f, 0x4f, 0xca, 0xdf, 0x7f, 0xfd, 0xfe, 0xe1, 0x95,
	0x48, 0xc1, 0xee, 0xef, 0xe2, 0xb8, 0x21, 0xf9, 0xf2, 0x04, 0xd5, 0xf0, 0x7b, 0xc8, 0x77, 0xb8,
	0x76, 0x0b, 0xb4, 0xbf, 0xf5, 0xe6, 0x49, 0xe8, 0xd6, 0x30, 0x48, 0xc5, 0xc6, 0xed, 0x61, 0x9c,
	0xc6, 0xd9, 0xe1, 0x34, 0xbe, 0x8a, 0xd1, 0x37, 0x7c, 0x05, 0x59, 0x33, 0x73, 0xbc, 0x89, 0xb2,
	0xb1, 0x45, 0x95, 0xfd, 0x2d, 0x3d, 0x61, 0x3c, 0xb0, 0xa1, 0x65, 0xb2, 0x9b, 0x86, 0x4e, 0xd5,
	0xc2, 0x65, 0x9e, 0xa0, 0x5a, 0xcb, 0xbf, 0xce, 0xb0, 0x99, 0x18, 0xe4, 0xec, 0xff, 0xf6, 0xf2,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x52, 0x50, 0x09, 0xab, 0x03, 0x00, 0x00,
}