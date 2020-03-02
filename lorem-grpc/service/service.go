package service

import (
	"context"
)

// Service ...
type Service interface {
	Lorem(ctx context.Context, requestType string, min, max int) (string, error)
}
