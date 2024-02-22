package elasticsearch

import (
	"context"
	"io"
)

// Client defines the interface for our Elasticsearch operations.
type Client interface {
	CreateDocument(ctx context.Context, index string, documentID string, body io.Reader) (interface{}, error)
	GetDocument(ctx context.Context, index string, documentID string) (interface{}, error)
	UpdateDocument(ctx context.Context, index string, documentID string, body io.Reader) (interface{}, error)
	DeleteDocument(ctx context.Context, index string, documentID string) (interface{}, error)
	Search(ctx context.Context, index string, body io.Reader) (interface{}, error)
	CreateIndex(ctx context.Context, index string, body io.Reader) (interface{}, error)
	GetIndex(ctx context.Context, index string) (interface{}, error)
}
