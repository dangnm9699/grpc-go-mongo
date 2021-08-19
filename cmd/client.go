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
	"github.com/dangnm9699/grpc-example/logger"
	pb "github.com/dangnm9699/grpc-example/pkg/movie"
	"google.golang.org/grpc"
	"time"

	"github.com/spf13/cobra"
)

const (
	ServerAddress = "localhost:5000"

	ClientTimeout = 1 * time.Second
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "gRPC Client",
	Long:  `gRPC Client`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
		if err != nil {
			logger.Error().Printf("cannot connect: %v\n", err)
		}
		defer conn.Close()
		client := pb.NewMovieSvcClient(conn)
		c, cancel := context.WithTimeout(context.Background(), ClientTimeout)
		defer cancel()
		r, err := client.PutMovie(c, &pb.PutMovieRequest{
			Movie: &pb.Movie{
				Tconst:      "tt4154756",
				Name:        "Avengers: Infinity War",
				ReleaseDate: "2018-04-25",
				Country:     "United States",
				Runtime:     149,
				MpaRating:   "PG-13",
			},
		})
		if err != nil {
			logger.Error().Fatalf("error occurred: %v", err)
		}
		logger.Debug().Printf("code=%d, message=%s\n", r.StatusCode, r.Message)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
