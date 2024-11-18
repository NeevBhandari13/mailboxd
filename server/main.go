package main

import (
	"context"
	"database/sql"
	"github.com/NeevB13/grpc-todo/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
	"strconv"
)

const (
	port = "localhost:9090"
)

var catalogue []*protos.MovieInfo

type movieServer struct {
	protos.UnimplementedMovieServer
}

func main() {
	initMovies()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	protos.RegisterMovieServer(s, &movieServer{})

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initMovies() {
	catalogue = append(catalogue, &protos.MovieInfo{Id: "1", Title: "Fight Club", Genre: protos.Genre_ACTION, Director: &protos.Director{FirstName: "David", LastName: "Fincher"}})
	catalogue = append(catalogue, &protos.MovieInfo{Id: "2", Title: "Spider-Man", Genre: protos.Genre_ACTION, Director: &protos.Director{FirstName: "Sam", LastName: "Raimi"}})
	catalogue = append(catalogue, &protos.MovieInfo{Id: "3", Title: "Superbad", Genre: protos.Genre_COMEDY, Director: &protos.Director{FirstName: "Greg", LastName: "Mottola"}})
}

func (s *movieServer) GetMovies(in *protos.Empty, stream protos.Movie_GetMoviesServer) error {
	log.Printf("Received: %v", in)
	for _, movie := range catalogue {
		if err := stream.Send(movie); err != nil {
			log.Fatalf("failed to send movie: %v", err)
			return err
		}
	}
	return nil
}

func (s *movieServer) GetMovie(ctx context.Context, id *protos.Id) (*protos.MovieInfo, error) {
	log.Printf("Received GetMovie request: %v", id)

	for _, movie := range catalogue {
		if movie.GetId() == id.GetValue() {
			return movie, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Movie not found")
}

func (s *movieServer) CreateMovie(ctx context.Context, movie *protos.MovieInfo) (*protos.Id, error) {
	log.Printf("Received CreateMovie request: %v", movie)

	movie.Id = strconv.Itoa(rand.Intn(1000))
	catalogue = append(catalogue, movie)
	return &protos.Id{Value: movie.GetId()}, nil
}

func (s *movieServer) UpdateMovie(ctx context.Context, movie *protos.MovieInfo) (*protos.MovieInfo, error) {
	log.Printf("Received UpdateMovie request: %v", movie)

	for i, m := range catalogue {
		if m.GetId() == movie.GetId() {
			catalogue[i] = movie
			return movie, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Movie not found")
}

func (s *movieServer) DeleteMovie(ctx context.Context, id *protos.Id) (*protos.MovieInfo, error) {
	log.Printf("Received DeleteMovie request: %v", id)

	for i, movie := range catalogue {
		if movie.GetId() == id.GetValue() {
			res := catalogue[i]
			catalogue = append(catalogue[:i], catalogue[i+1:]...)
			return res, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Movie not found")
}
