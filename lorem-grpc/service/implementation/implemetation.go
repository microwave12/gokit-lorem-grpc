package implementation

import (
	"context"
	"errors"
	"strings"

	lorem "github.com/drhodes/golorem"
)

var (
	// ErrRequestTypeNotFound ...
	ErrRequestTypeNotFound = errors.New("request type only valid for word, sentence, paragraph")
)

// LoremService ...
type LoremService struct{}

// Lorem ...
func (LoremService) Lorem(_ context.Context, requestType string, min, max int) (string, error) {
	var result string
	var err error

	if strings.EqualFold(requestType, "Word") {
		result = lorem.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = lorem.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = lorem.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}

	return result, err
}
