package service

import "context"

// Service ...
type Service interface {
	Word(ctx context.Context, min, max int) (string, error)
	Sentence(ctx context.Context, min, max int) (string, error)
	Paragraph(ctx context.Context, min, max int) (string, error)
}
