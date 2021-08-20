// Package cmd includes CLI commands
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
	"github.com/dangnm9699/grpc-go-mongo/logger"
	pb "github.com/dangnm9699/grpc-go-mongo/pkg/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/spf13/cobra"
)

// Server implements movie service server
type Server struct {
	pb.UnimplementedMovieSvcServer
}

const (
	// ServerPort is gRPC server port
	ServerPort = ":5000"

	// MongoUri are used to connect to MongoDB
	MongoUri = "mongodb://localhost:27017"
	// MongoDatabase specifies database to connect
	MongoDatabase = "grpc"
	// MongoCollection specifies collection to connect
	MongoCollection = "example"
	// MongoTimeout represents Mongo timeout duration
	MongoTimeout = 10 * time.Second
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

// PutMovie method create a movie or update a existed movie
func (svr *Server) PutMovie(ctx context.Context, req *pb.PutMovieRequest) (*pb.PutMovieResponse, error) {
	var movie *pb.Movie
	movie = req.GetMovie()
	if len(movie.Tconst) == 0 || len(movie.Name) == 0 {
		logger.Debug().Printf("put movie failed: null values\n")
		return nil, fmt.Errorf("null values")
	}
	ur, err := mongoColl.ReplaceOne(
		ctx,
		bson.M{"tconst": movie.Tconst},
		movie,
		&options.ReplaceOptions{
			Upsert: &mongoUpsert,
		},
	)
	if err != nil {
		logger.Error().Printf("put movie{tconst=%s} failed: %v\n", movie.Tconst, err)
		return nil, fmt.Errorf("put movie{tconst=%s}: database failed", movie.Tconst)
	}
	logger.Debug().Printf("put movie{tconst=%s} successfully\n", movie.Tconst)
	if ur.UpsertedCount == 0 {
		return &pb.PutMovieResponse{
			Message: fmt.Sprintf("put movie{tconst=%s}: updated", movie.Tconst),
		}, nil
	}
	return &pb.PutMovieResponse{
		Message: fmt.Sprintf("put movie{tconst=%s}: created", movie.Tconst),
	}, nil
}

// GetMovie method retrieve a specified movie if exists
func (svr *Server) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	tconst := req.GetTconst()
	if len(tconst) == 0 {
		logger.Error().Printf("get movie failed: null values\n")
		return nil, fmt.Errorf("null values")
	}
	sr := mongoColl.FindOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if sr.Err() == mongo.ErrNoDocuments {
		logger.Error().Printf("get movie{tconst=%s} failed: %v\n", tconst, sr.Err())
		return nil, fmt.Errorf("get movie{tconst=%s}: not found", tconst)
	}
	var movie *pb.Movie
	err := sr.Decode(movie)
	if err != nil {
		logger.Error().Printf("get movie{tconst=%s} failed: %v\n", tconst, err)
		return nil, fmt.Errorf("get movie{tconst=%s}: decode failed", tconst)
	}
	logger.Debug().Printf("get movie{tconst=%s} successfully\n", tconst)
	return &pb.GetMovieResponse{
		Movie: movie,
	}, nil
}

// GetMovies method retrieve all movies
func (svr *Server) GetMovies(ctx context.Context, _ *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {
	cur, err := mongoColl.Find(
		ctx,
		bson.M{},
	)
	if err != nil {
		logger.Error().Printf("get movies failed: %v\n", err)
		return nil, fmt.Errorf("get movies: database failed")
	}
	var movies []*pb.Movie
	for cur.Next(ctx) {
		var movie *pb.Movie
		err := cur.Decode(movie)
		if err != nil {
			logger.Error().Printf("get movies failed: %v\n", err)
			movies = nil
			return nil, fmt.Errorf("get movies: retrieve database failed")
		}
		movies = append(movies, movie)
	}
	logger.Debug().Printf("get movies successfully\n")
	return &pb.GetMoviesResponse{
		Message: "get movies: ok",
		Movies:  movies,
	}, nil
}

// DeleteMovie method delete a specified movie if exists
func (svr *Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	tconst := req.GetTconst()
	if len(tconst) == 0 {
		logger.Error().Printf("delete movie failed: null values\n")
		return nil, fmt.Errorf("null values")
	}
	dr, err := mongoColl.DeleteOne(
		ctx,
		bson.M{"tconst": tconst},
	)
	if err != nil {
		logger.Error().Printf("delete movie{tconst=%s} failed: %v\n", tconst, err)
		return nil, fmt.Errorf("delete movie{tconst=%s}: database failed", tconst)
	}
	if dr.DeletedCount == 0 {
		logger.Debug().Printf("delete movie{tconst=%s} failed: not available\n", tconst)
		return nil, fmt.Errorf("delete movie{tconst=%s}: not found", tconst)
	}
	logger.Debug().Printf("delete movie{tconst=%s} successfully\n", tconst)
	return &pb.DeleteMovieResponse{
		Message: fmt.Sprintf("delete movie{tconst=%s}: ok", tconst),
	}, nil
}
