// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: protos/movies.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Movie_GetMovies_FullMethodName   = "/movies.Movie/GetMovies"
	Movie_GetMovie_FullMethodName    = "/movies.Movie/GetMovie"
	Movie_CreateMovie_FullMethodName = "/movies.Movie/CreateMovie"
	Movie_UpdateMovie_FullMethodName = "/movies.Movie/UpdateMovie"
	Movie_DeleteMovie_FullMethodName = "/movies.Movie/DeleteMovie"
)

// MovieClient is the client API for Movie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieClient interface {
	GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Movie_GetMoviesClient, error)
	GetMovie(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MovieInfo, error)
	CreateMovie(ctx context.Context, in *MovieInfo, opts ...grpc.CallOption) (*Id, error)
	UpdateMovie(ctx context.Context, in *MovieInfo, opts ...grpc.CallOption) (*MovieInfo, error)
	DeleteMovie(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MovieInfo, error)
}

type movieClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieClient(cc grpc.ClientConnInterface) MovieClient {
	return &movieClient{cc}
}

func (c *movieClient) GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Movie_GetMoviesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Movie_ServiceDesc.Streams[0], Movie_GetMovies_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &movieGetMoviesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Movie_GetMoviesClient interface {
	Recv() (*MovieInfo, error)
	grpc.ClientStream
}

type movieGetMoviesClient struct {
	grpc.ClientStream
}

func (x *movieGetMoviesClient) Recv() (*MovieInfo, error) {
	m := new(MovieInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *movieClient) GetMovie(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MovieInfo, error) {
	out := new(MovieInfo)
	err := c.cc.Invoke(ctx, Movie_GetMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) CreateMovie(ctx context.Context, in *MovieInfo, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, Movie_CreateMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) UpdateMovie(ctx context.Context, in *MovieInfo, opts ...grpc.CallOption) (*MovieInfo, error) {
	out := new(MovieInfo)
	err := c.cc.Invoke(ctx, Movie_UpdateMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) DeleteMovie(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MovieInfo, error) {
	out := new(MovieInfo)
	err := c.cc.Invoke(ctx, Movie_DeleteMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServer is the server API for Movie service.
// All implementations must embed UnimplementedMovieServer
// for forward compatibility
type MovieServer interface {
	GetMovies(*Empty, Movie_GetMoviesServer) error
	GetMovie(context.Context, *Id) (*MovieInfo, error)
	CreateMovie(context.Context, *MovieInfo) (*Id, error)
	UpdateMovie(context.Context, *MovieInfo) (*MovieInfo, error)
	DeleteMovie(context.Context, *Id) (*MovieInfo, error)
	mustEmbedUnimplementedMovieServer()
}

// UnimplementedMovieServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServer struct {
}

func (UnimplementedMovieServer) GetMovies(*Empty, Movie_GetMoviesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieServer) GetMovie(context.Context, *Id) (*MovieInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieServer) CreateMovie(context.Context, *MovieInfo) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedMovieServer) UpdateMovie(context.Context, *MovieInfo) (*MovieInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieServer) DeleteMovie(context.Context, *Id) (*MovieInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieServer) mustEmbedUnimplementedMovieServer() {}

// UnsafeMovieServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServer will
// result in compilation errors.
type UnsafeMovieServer interface {
	mustEmbedUnimplementedMovieServer()
}

func RegisterMovieServer(s grpc.ServiceRegistrar, srv MovieServer) {
	s.RegisterService(&Movie_ServiceDesc, srv)
}

func _Movie_GetMovies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieServer).GetMovies(m, &movieGetMoviesServer{stream})
}

type Movie_GetMoviesServer interface {
	Send(*MovieInfo) error
	grpc.ServerStream
}

type movieGetMoviesServer struct {
	grpc.ServerStream
}

func (x *movieGetMoviesServer) Send(m *MovieInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Movie_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Movie_GetMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).GetMovie(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Movie_CreateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).CreateMovie(ctx, req.(*MovieInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Movie_UpdateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).UpdateMovie(ctx, req.(*MovieInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Movie_DeleteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).DeleteMovie(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Movie_ServiceDesc is the grpc.ServiceDesc for Movie service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Movie_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movies.Movie",
	HandlerType: (*MovieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _Movie_GetMovie_Handler,
		},
		{
			MethodName: "CreateMovie",
			Handler:    _Movie_CreateMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _Movie_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _Movie_DeleteMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMovies",
			Handler:       _Movie_GetMovies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/movies.proto",
}
