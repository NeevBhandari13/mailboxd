package main

import (
	"context"
	"github.com/NeevB13/grpc-todo/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:9000"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := protos.NewMovieClient(conn)
	//
	//callGetMovies(ctx, client)
	//
	//callGetMovie(ctx, client, "1")
	//
	//newMovie := &protos.MovieInfo{
	//	Title:    "Interstellar",
	//	Genre:    protos.Genre_SCI_FI,
	//	Director: &protos.Director{FirstName: "Christopher", LastName: "Nolan"},
	//}
	//callCreateMovie(ctx, client, newMovie)
	//
	//// Example call to DeleteMovie
	//callDeleteMovie(ctx, client, "1")

}

func callGetMovies(ctx context.Context, client protos.MovieClient) {

	req := &protos.Empty{}
	stream, err := client.GetMovies(ctx, req)
	if err != nil {
		log.Fatalf("Error from getMovies: %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error from getMovies stream: %v", err)
		}
		log.Printf("Movie: %v", row)
	}
}

func callGetMovie(ctx context.Context, client protos.MovieClient, id string) {

	movie, err := client.GetMovie(ctx, &protos.Id{Value: id})
	if err != nil {
		log.Fatalf("could not get movie: %v", err)
	}
	log.Printf("Movie: %v", movie)
}

func callCreateMovie(ctx context.Context, client protos.MovieClient, movie *protos.MovieInfo) {
	id, err := client.CreateMovie(ctx, movie)
	if err != nil {
		log.Fatalf("could not create movie: %v", err)
	}
	log.Printf("Created movie with ID: %v", id.Value)
}

func callDeleteMovie(ctx context.Context, client protos.MovieClient, id string) {
	deletedMovie, err := client.DeleteMovie(ctx, &protos.Id{Value: id})
	if err != nil {
		log.Fatalf("could not delete movie: %v", err)
	}
	log.Printf("Deleted movie: %v", deletedMovie)
}
