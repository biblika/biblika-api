package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/biblika/biblika-api/api"
	"github.com/biblika/biblika-api/proto/book"
	"github.com/biblika/biblika-api/proto/passage"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	r := runtime.NewServeMux()
	err := book.RegisterBookServiceHandlerServer(context.Background(), r, &api.BookService{})
	if err != nil {
		log.Fatalf("failed to register book service: %v", err)
	}

	err = passage.RegisterPassageServiceHandlerServer(context.Background(), r, &api.PassageService{})
	if err != nil {
		log.Fatalf("failed to register passage service: %v", err)
	}

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = ":8080"
	}

	log.Fatal(http.ListenAndServe(port, r))
}
