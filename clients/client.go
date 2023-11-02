package client

import (
	"context"
	"net/url"

	"github.com/rohanraj7316/httpclient"
)

const (
	ENV_API_BASE_URL = "https://www.alphavantage.co/query"
	ENV_API_KEY      = "5OIALBMMNV7CBIPD"
	// ENV_API_KEY      = "PQ9CJR119JHM1ZUP"
)

type Stocks interface {
	IntraDay(ctx context.Context, symbol, month string) (*IntraDayResponse, error)
}

type Intelligence interface {
	NewsSentiments(ctx context.Context, symbol, to, from string) (*NewsSentimentsResponse, error)
}

type alphaVantage struct {
	client  *httpclient.HttpClient
	baseUrl *url.URL
}

func NewStock() (Stocks, error) {
	cl, err := httpclient.New()
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(ENV_API_BASE_URL)
	if err != nil {
		return nil, err
	}

	return &alphaVantage{
		client:  cl,
		baseUrl: u,
	}, nil
}

func NewIntelligence() (Intelligence, error) {
	cl, err := httpclient.New()
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(ENV_API_BASE_URL)
	if err != nil {
		return nil, err
	}

	return &alphaVantage{
		client:  cl,
		baseUrl: u,
	}, nil
}
