package client

import (
	"context"
)

// Datasets describes all the dataset-related methods that the Honeycomb API
// supports.
//
// API docs: https://docs.honeycomb.io/api/datasets/
type Datasets interface {
	// List all datasets.
	List(ctx context.Context) ([]Dataset, error)

	// Get a dataset by its slug. Returns ErrNotFound if there is no dataset
	// with the given slug.
	Get(ctx context.Context, slug string) (*Dataset, error)

	// Create a new dataset. Only name should be set when creating a dataset,
	// all other fields are ignored.
	Create(ctx context.Context, dataset *Dataset) (*Dataset, error)
}

// datasets implements Datasets.
type datasets struct {
	client *Client
}

// Compile-time proof of interface implementation by type datasets.
var _ Datasets = (*datasets)(nil)

// Dataset represents a Honeycomb dataset.
//
// API docs: https://docs.honeycomb.io/api/dataset
type Dataset struct {
	Name string `json:"name"`
	Slug string `json:"slug,omitempty"`
}

func (s datasets) List(ctx context.Context) ([]Dataset, error) {
	var datasets []Dataset
	err := s.client.performRequest(ctx, "GET", "/1/datasets", nil, &datasets)
	return datasets, err
}

func (s datasets) Get(ctx context.Context, slug string) (*Dataset, error) {
	var dataset Dataset
	err := s.client.performRequest(ctx, "GET", "/1/datasets/"+slug, nil, &dataset)
	return &dataset, err
}

func (s datasets) Create(ctx context.Context, data *Dataset) (*Dataset, error) {
	var dataset Dataset
	err := s.client.performRequest(ctx, "POST", "/1/datasets", data, &dataset)
	return &dataset, err
}
