package main

import (
	"context"
	"fmt"
	pb "go-grpc-client/generated"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

// Your gRPC server addresses
const (
	greeterAddress   = "localhost:50051" // replace with your actual address
	bookStoreAddress = "localhost:50051" // replace with your actual address
	movieClubAddress = "localhost:50051" // replace with your actual address
)

func main() {
	// Create gRPC connections to each service
	greeterConn, err := grpc.Dial(greeterAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Greeter service: %v", err)
	}
	defer greeterConn.Close()

	bookStoreConn, err := grpc.Dial(bookStoreAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to BookStore service: %v", err)
	}
	defer bookStoreConn.Close()

	movieClubConn, err := grpc.Dial(movieClubAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to MovieClub service: %v", err)
	}
	defer movieClubConn.Close()

	// Create gRPC clients for each service
	greeterClient := pb.NewGreeterClient(greeterConn)
	bookStoreClient := pb.NewBookStoreClient(bookStoreConn)
	movieClubClient := pb.NewMovieClubClient(movieClubConn)

	// Define a handler function for the web server
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// Make a gRPC call to Greeter service
		greeterResponse, err := greeterClient.Greet(context.Background(), &pb.ClientInput{Greeting: "Hello", Name: "User"})
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call Greeter service: %v", err), http.StatusInternalServerError)
			return
		}

		// Display the response from Greeter service
		fmt.Fprintf(w, "Greeter Response: %s\n", greeterResponse.Message)
	})

	http.HandleFunc("/search/book", func(w http.ResponseWriter, r *http.Request) {
		// Extract book search parameters from the query string
		query := r.URL.Query()
		bookName := query.Get("name")
		bookAuthor := query.Get("author")
		bookGenre := query.Get("genre")

		// Make a gRPC call to BookStore service
		bookResponse, err := bookStoreClient.First(context.Background(), &pb.BookSearch{Name: bookName, Author: bookAuthor, Genre: bookGenre})
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call BookStore service: %v", err), http.StatusInternalServerError)
			return
		}

		// Display the response from BookStore service
		fmt.Fprintf(w, "Book Response: %+v\n", bookResponse)
	})

	http.HandleFunc("/search/movie", func(w http.ResponseWriter, r *http.Request) {
		// Extract movie search parameters from the query string
		query := r.URL.Query()
		movieName := query.Get("name")
		movieDirector := query.Get("director")
		movieGenre := query.Get("genre")

		// Make a gRPC call to MovieClub service
		movieResponse, err := movieClubClient.First(context.Background(), &pb.MovieSearch{Name: movieName, Director: movieDirector, Genre: movieGenre})
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call MovieClub service: %v", err), http.StatusInternalServerError)
			return
		}

		// Display the response from MovieClub service
		fmt.Fprintf(w, "Movie Response: %+v\n", movieResponse)
	})

	// Start the web server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
