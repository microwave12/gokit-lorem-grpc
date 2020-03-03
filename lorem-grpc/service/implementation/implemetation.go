package implementation

import (
	"context"
	"errors"

	lorem "github.com/drhodes/golorem"
)

var (
	// ErrRequestTypeNotFound ...
	ErrRequestTypeNotFound = errors.New("request type only valid for word, sentence, paragraph")
)

// LoremService ...
type LoremService struct{}

// Word ...
func (LoremService) Word(ctx context.Context, min, max int) (string, error) {
	return lorem.Word(min, max), nil
}

// Sentence ...
func (LoremService) Sentence(ctx context.Context, min, max int) (string, error) {
	return lorem.Sentence(min, max), nil
}

// Paragraph ...
func (LoremService) Paragraph(ctx context.Context, min, max int) (string, error) {
	return lorem.Paragraph(min, max), nil
}
