package main

import (
	"context"
	"fmt"
	"github.com/dangnm9699/grpc-example/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

type MovieSvc struct {
	pb.UnimplementedMovieSvcServer
}

func (ms *MovieSvc) RunGRPC(c context.Context) error {
	ltn, err := net.Listen("tcp", ":"+config.port)
	if err != nil {
		logger.Error(err, fmt.Sprintf("cannot create tcp port:%s", config.port))
		return err
	}
	server := grpc.NewServer()
	pb.RegisterMovieSvcServer(server, ms)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		logger.Info("shutting down gRPC server...")
		server.GracefulStop()
		<-ctx.Done()
	}()
	logger.Info("starting down gRPC server...")
	return server.Serve(ltn)
}

func (ms *MovieSvc) PutMovie(c context.Context, req *pb.PutMovieRequest) (*pb.PutMovieResponse, error) {
	mov := req.GetMovie()
	_, err := mongodb.coll.ReplaceOne(
		ctx,
		bson.M{"tconst": mov.Tconst},
		mov,
		&options.ReplaceOptions{
			Upsert: &upsert,
		},
	)
	if err != nil {
		logger.Error(err, fmt.Sprintf("put movie{tconst=%s} failed", mov.Tconst))
		return &pb.PutMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
		}, err
	}
	logger.Info(fmt.Sprintf("put movie{tconst=%s} successfully", mov.Tconst))
	return &pb.PutMovieResponse{
		StatusCode: 200,
		Message:    "ok",
	}, nil
}

func (ms *MovieSvc) GetMovie(c context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	tconst := req.GetTconst()
	sr := mongodb.coll.FindOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if sr.Err() == mongo.ErrNoDocuments {
		logger.Error(sr.Err(), fmt.Sprintf("get movie{tconst=%s} failed, not found", tconst))
		return &pb.GetMovieResponse{
			StatusCode: 404,
			Message:    "not found",
			Movie:      nil,
		}, fmt.Errorf("movie not found")
	}
	var mov *pb.Movie
	err := sr.Decode(mov)
	if err != nil {
		logger.Error(err, fmt.Sprintf("get movie{tconst=%s} failed, decoded failed", tconst))
		return &pb.GetMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
			Movie:      nil,
		}, err
	}
	logger.Info(fmt.Sprintf("get movie{tconst=%s} successfully", tconst))
	return &pb.GetMovieResponse{
		StatusCode: 200,
		Message:    "ok",
		Movie:      mov,
	}, nil
}

func (ms *MovieSvc) GetMovies(c context.Context, req *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {
	cur, err := mongodb.coll.Find(
		ctx,
		bson.M{},
	)
	if err != nil {
		logger.Error(err, fmt.Sprintf("get movies failed, mongodb error"))
		return &pb.GetMoviesResponse{
			StatusCode: 500,
			Message:    "internal server error",
			Movies:     nil,
		}, err
	}
	var movs []*pb.Movie
	for cur.Next(ctx) {
		var mov *pb.Movie
		err := cur.Decode(mov)
		if err != nil {
			logger.Error(err, fmt.Sprintf("get movies failed, decode error"))
			movs = nil
			return &pb.GetMoviesResponse{
				StatusCode: 500,
				Message:    "internal server error",
				Movies:     nil,
			}, err
		}
		movs = append(movs, mov)
	}
	logger.Info(fmt.Sprintf("get movies successfully"))
	return &pb.GetMoviesResponse{
		StatusCode: 200,
		Message:    "ok",
		Movies:     movs,
	}, nil
}

func (ms *MovieSvc) DeleteMovie(c context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	tconst := req.GetTconst()
	dr, err := mongodb.coll.DeleteOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if err != nil {
		logger.Error(err, fmt.Sprintf("delete movie{tconst=%s} failed, mongodb error", tconst))
		return &pb.DeleteMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
		}, err
	}
	if dr.DeletedCount == 0 {
		logger.Error(err, fmt.Sprintf("delete movie{tconst=%s} failed, not found", tconst))
		return &pb.DeleteMovieResponse{
			StatusCode: 404,
			Message:    "not found",
		}, fmt.Errorf("movie not found")
	}
	logger.Info(fmt.Sprintf("delete movie{tconst=%s} successfully", tconst))
	return &pb.DeleteMovieResponse{
		StatusCode: 200,
		Message:    "ok",
	}, nil
}
