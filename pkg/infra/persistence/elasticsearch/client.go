package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Audit struct {
	date      time.Time
	operation string
}

var _ Client = &esClient{}

type esClient struct {
	client *elasticsearch.Client
}

// NewClient creates a new Elasticsearch client instance
func NewClient(config elasticsearch.Config) (Client, error) {
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &esClient{client: es}, nil
}

// CreateDocument creates a new document
func (e *esClient) CreateDocument(ctx context.Context, index string, documentID string, body io.Reader) (interface{}, error) {
	res, err := e.client.Index(index, body, e.client.Index.WithDocumentID(documentID), e.client.Index.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// GetDocument retrieves a document with the specified ID
func (e *esClient) GetDocument(ctx context.Context, index string, documentID string) (interface{}, error) {
	res, err := e.client.Get(index, documentID, e.client.Get.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// UpdateDocument updates a document with the specified ID
func (e *esClient) UpdateDocument(ctx context.Context, index string, documentID string, body io.Reader) (interface{}, error) {
	req := map[string]interface{}{
		"doc": body,
	}
	updateBody, _ := json.Marshal(req)
	res, err := e.client.Update(index, documentID, bytes.NewReader(updateBody), e.client.Update.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// DeleteDocument deletes a document with the specified ID
func (e *esClient) DeleteDocument(ctx context.Context, index string, documentID string) (interface{}, error) {
	res, err := e.client.Delete(index, documentID, e.client.Delete.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// Search performs a search query in the specified index
func (e *esClient) Search(ctx context.Context, index string, body io.Reader) (interface{}, error) {
	var buf bytes.Buffer
	io.Copy(&buf, body)
	res, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// parseResponse parses the response body and returns the result
func parseResponse(res *esapi.Response) (interface{}, error) {
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("error parsing the response body: %s", err)
		} else {
			// print out the detailed information of the request and response
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}
	return r, nil
}

// CreateIndex creates a new index with the specified settings and mappings
func (e *esClient) CreateIndex(ctx context.Context, index string, body io.Reader) (interface{}, error) {
	res, err := e.client.Indices.Create(index, e.client.Indices.Create.WithBody(body), e.client.Indices.Create.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

// GetIndex retrieves the settings and mappings of an index
func (e *esClient) GetIndex(ctx context.Context, index string) (interface{}, error) {
	res, err := e.client.Indices.Get([]string{index}, e.client.Indices.Get.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}
