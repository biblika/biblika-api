package api

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/biblika/biblika-api/proto/book"
	"github.com/biblika/biblika-api/util"
)

type BookService struct {
	book.UnimplementedBookServiceServer
}

var (
	bookResponseOnce sync.Once
	byteOnce         []byte
	errOnce          error
)

func (s *BookService) Book(ctx context.Context, req *book.BookRequest) (*book.BookResponse, error) {
	bookResponseOnce.Do(func() {
		reader := util.NewFileReader("./json/book.json")
		byteOnce, errOnce = reader.JSON()
	})

	if errOnce != nil {
		log.Fatalf("failed to read file: %v", errOnce)
	}

	var response book.BookResponse
	err := json.Unmarshal(byteOnce, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
