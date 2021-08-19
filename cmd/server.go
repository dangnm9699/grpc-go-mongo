// Package cmd
/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/dangnm9699/grpc-example/logger"
	pb "github.com/dangnm9699/grpc-example/pkg/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/spf13/cobra"
)

type Server struct {
	pb.UnimplementedMovieSvcServer
}

const (
	ServerPort = ":5000"

	MongoUri        = "mongodb://localhost:27017"
	MongoDatabase   = "grpc"
	MongoCollection = "example"
	MongoTimeout    = 10 * time.Second
)

var mongoColl *mongo.Collection
var mongoUpsert = true

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "gRPC server",
	Long:  `gRPC server connect to MongoDB`,
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to MongoDB
		c, cancel := context.WithTimeout(context.Background(), MongoTimeout)
		defer cancel()
		client, err := mongo.Connect(c, options.Client().ApplyURI(MongoUri))
		if err != nil {
			logger.Error().Fatalf("cannot connect to mongodb: %v\n", err)
		}
		logger.Info().Printf("successfully connect to mongodb\n")
		mongoColl = client.Database(MongoDatabase).Collection(MongoCollection)
		// Start server
		lis, err := net.Listen("tcp", ServerPort)
		if err != nil {
			logger.Error().Fatalf("failed to listen: %v\n", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterMovieSvcServer(grpcServer, &Server{})
		logger.Debug().Printf("gRPC server listening on %v\n", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error().Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func (svr *Server) PutMovie(ctx context.Context, req *pb.PutMovieRequest) (*pb.PutMovieResponse, error) {
	mov := req.GetMovie()
	_, err := mongoColl.ReplaceOne(
		ctx,
		bson.M{"tconst": mov.Tconst},
		mov,
		&options.ReplaceOptions{
			Upsert: &mongoUpsert,
		},
	)
	if err != nil {
		logger.Error().Printf("put movie{tconst=%s} failed: %v\n", mov.Tconst, err)
		return &pb.PutMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
		}, err
	}
	logger.Debug().Printf("put movie{tconst=%s} successfully\n", mov.Tconst)
	return &pb.PutMovieResponse{
		StatusCode: 200,
		Message:    "ok",
	}, nil
}

func (svr *Server) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	tconst := req.GetTconst()
	sr := mongoColl.FindOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if sr.Err() == mongo.ErrNoDocuments {
		logger.Error().Printf("get movie{tconst=%s} failed: %v\n", tconst, sr.Err())
		return &pb.GetMovieResponse{
			StatusCode: 404,
			Message:    "not found",
			Movie:      nil,
		}, fmt.Errorf("movie not found")
	}
	var mov *pb.Movie
	err := sr.Decode(mov)
	if err != nil {
		logger.Error().Printf("get movie{tconst=%s} failed: %v\n", tconst, err)
		return &pb.GetMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
			Movie:      nil,
		}, err
	}
	logger.Debug().Printf("get movie{tconst=%s} successfully\n", tconst)
	return &pb.GetMovieResponse{
		StatusCode: 200,
		Message:    "ok",
		Movie:      mov,
	}, nil
}

func (svr *Server) GetMovies(ctx context.Context, _ *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {
	cur, err := mongoColl.Find(
		ctx,
		bson.M{},
	)
	if err != nil {
		logger.Error().Printf("get movies failed: %v\n", err)
		return &pb.GetMoviesResponse{
			StatusCode: 500,
			Message:    "internal server error",
			Movies:     nil,
		}, err
	}
	var movies []*pb.Movie
	for cur.Next(ctx) {
		var movie *pb.Movie
		err := cur.Decode(movie)
		if err != nil {
			logger.Error().Printf("get movies failed: %v\n", err)
			movies = nil
			return &pb.GetMoviesResponse{
				StatusCode: 500,
				Message:    "internal server error",
				Movies:     nil,
			}, err
		}
		movies = append(movies, movie)
	}
	logger.Debug().Printf("get movies successfully\n")
	return &pb.GetMoviesResponse{
		StatusCode: 200,
		Message:    "ok",
		Movies:     movies,
	}, nil
}

func (svr *Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	tconst := req.GetTconst()
	dr, err := mongoColl.DeleteOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if err != nil {
		logger.Error().Printf("delete movie{tconst=%s} failed: %v\n", tconst, err)
		return &pb.DeleteMovieResponse{
			StatusCode: 500,
			Message:    "internal server error",
		}, err
	}
	if dr.DeletedCount == 0 {
		logger.Debug().Printf("delete movie{tconst=%s} failed: not available\n", tconst)
		return &pb.DeleteMovieResponse{
			StatusCode: 404,
			Message:    "not found",
		}, fmt.Errorf("movie not found")
	}
	logger.Debug().Printf("delete movie{tconst=%s} successfully\n", tconst)
	return &pb.DeleteMovieResponse{
		StatusCode: 200,
		Message:    "ok",
	}, nil
}
