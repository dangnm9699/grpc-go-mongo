package cmd

import (
	"context"
	"fmt"
	"github.com/dangnm9699/grpc-go-mongo/logger"
	pb "github.com/dangnm9699/grpc-go-mongo/pkg/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func init() {
	c, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	client, err := mongo.Connect(c, options.Client().ApplyURI(MongoUri))
	if err != nil {
		logger.Error().Fatalf("cannot connect to mongodb: %v\n", err)
	}
	logger.Info().Printf("successfully connect to mongodb\n")
	mongoColl = client.Database(MongoDatabase).Collection(MongoCollection)
	_, _ = mongoColl.DeleteOne(c, bson.M{"tconst": "tt4154796"})
}

func TestServer_PutMovie(t *testing.T) {
	server := &Server{}

	testCases := []struct {
		Name            string
		Request         *pb.PutMovieRequest
		ExpectedMessage string
		ExpectedError   bool
	}{
		{
			Name: "Put new movie",
			Request: &pb.PutMovieRequest{
				Movie: &pb.Movie{
					Tconst:      "tt4154796",
					Name:        "Avengers: Endgame",
					ReleaseDate: "2019-04-26",
					Country:     "United States",
					Runtime:     181,
					MpaRating:   "",
				},
			},
			ExpectedMessage: "put movie{tconst=tt4154796}: created",
			ExpectedError:   false,
		},
		{
			Name: "Put exited movie",
			Request: &pb.PutMovieRequest{
				Movie: &pb.Movie{
					Tconst:      "tt4154796",
					Name:        "Avengers: Endgame",
					ReleaseDate: "2019-04-26",
					Country:     "United States",
					Runtime:     181,
					MpaRating:   "PG-13",
				},
			},
			ExpectedMessage: "put movie{tconst=tt4154796}: updated",
			ExpectedError:   false,
		},
		{
			Name: "Put nil movie",
			Request: &pb.PutMovieRequest{
				Movie: &pb.Movie{},
			},
			ExpectedMessage: "",
			ExpectedError:   true,
		},
	}
	for _, tc := range testCases {
		testCase := tc
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			res, err := server.PutMovie(ctx, testCase.Request)
			if testCase.ExpectedError {
				if err.Error() != fmt.Sprintf("null values") {
					t.Errorf("PutMovie got error %s\n", err.Error())
				}
			} else {
				if res.Message != testCase.ExpectedMessage {
					t.Errorf("PutMovie{tconst=%s} got message = %s\n", testCase.Request.Movie.Tconst, res.Message)
				}
			}
		})
	}
}

func TestServer_GetMovie(t *testing.T) {

}

func TestServer_GetMovies(t *testing.T) {

}

func TestServer_DeleteMovie(t *testing.T) {

}
