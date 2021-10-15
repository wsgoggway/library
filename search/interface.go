package search

import (
	"context"
)

type SearchEngine interface {
	Search(ctx context.Context, searchString string, limit, offset int, index string, fields []string) (offersIds []int64, err error)
}
