package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// 20231031T224700
	CONST_NEWS_SENTIMENTAL_DATE_FORMAT = "20060102T150405"
)

type Feed struct {
	TimePublished         string  `json:"time_published,omitempty"`
	OverallSentimentLabel string  `json:"overall_sentiment_label,omitempty"`
	OverallSentimentScore float64 `json:"overall_sentiment_score,omitempty"`
}

type NewsSentimentsResponse struct {
	Feeds []Feed `json:"feed,omitempty"`
}

func (c alphaVantage) NewsSentiments(ctx context.Context, symbol, from, to string) (*NewsSentimentsResponse, error) {
	baseUrl, err := url.Parse(ENV_API_BASE_URL)
	if err != nil {
		return nil, err
	}

	q := baseUrl.Query()

	q.Add("function", "NEWS_SENTIMENT")
	q.Add("symbol", symbol)
	q.Add("time_from", from)
	q.Add("time_to", to)
	q.Add("sort", "LATEST")
	q.Add("apikey", ENV_API_KEY)
	q.Add("limit", "1000")

	baseUrl.RawQuery = q.Encode()

	r, err := c.client.Request(ctx, http.MethodGet, baseUrl.String(), map[string]string{}, nil)
	if err != nil {
		return nil, err
	}

	res := &NewsSentimentsResponse{}
	err = json.NewDecoder(r.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
