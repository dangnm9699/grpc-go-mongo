package main

import (
	"context"
	"github.com/dangnm9699/grpc-example/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

type MovieStub struct {
	clt pb.MovieSvcClient
}

func (ms *MovieStub) CreateOrModify(movie *pb.Movie) {
	req := &pb.PutMovieRequest{Movie: movie}
	resp, _ := ms.clt.PutMovie(ctx, req)
	log.Printf("status: %d ; msg: %s", resp.StatusCode, resp.Message)
}

func (ms *MovieStub) RetrieveOne(tconst string) *pb.Movie {
	req := &pb.GetMovieRequest{Tconst: tconst}
	resp, _ := ms.clt.GetMovie(ctx, req)
	log.Printf("status: %d ; msg: %s", resp.StatusCode, resp.Message)
	return resp.Movie
}

func (ms *MovieStub) RetrieveMany() []*pb.Movie {
	req := &pb.GetMoviesRequest{}
	resp, _ := ms.clt.GetMovies(ctx, req)
	log.Printf("status: %d ; msg: %s", resp.StatusCode, resp.Message)
	return resp.Movies
}

func (ms *MovieStub) Delete(tconst string) {
	req := &pb.DeleteMovieRequest{Tconst: tconst}
	resp, _ := ms.clt.DeleteMovie(ctx, req)
	log.Printf("status: %d ; msg: %s", resp.StatusCode, resp.Message)
}

var ctx = context.Background()

func main() {
	addr := "localhost:5000"
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot create connection to %s\n", addr)
	}
	defer conn.Close()
	movieClt := &MovieStub{pb.NewMovieSvcClient(conn)}
	log.Println("start sending requests")
	movie := &pb.Movie{
		Tconst:      "tt4154796",
		Name:        "Avengers: Endgame",
		ReleaseDate: time.Date(2019, time.April, 26, 0, 0, 0, 0, time.UTC).Unix(),
		Country:     "United States",
		Runtime:     181,
		MpaRating:   "C13",
	}
	movieClt.CreateOrModify(movie)
}
