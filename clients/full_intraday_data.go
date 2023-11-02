package client

import (
	"context"
	"encoding/json"
	"net/http"
)

const CONST_INTRA_DAY_DATE_FORMAT = "2006-01-02 15:04:05"

type Ticker struct {
	OneOpen    string `json:"1. open,omitempty"`
	TwoHigh    string `json:"2. high,omitempty"`
	ThreeLow   string `json:"3. low,omitempty"`
	FourClose  string `json:"4. close,omitempty"`
	FiveVolume string `json:"5. volume,omitempty"`
}

type IntraDayResponse struct {
	MetaData struct {
		OneInformation     string `json:"1. Information,omitempty"`
		TwoSymbol          string `json:"2. Symbol,omitempty"`
		ThreeLastRefreshed string `json:"3. Last Refreshed,omitempty"`
		FourInterval       string `json:"4. Interval,omitempty"`
		FiveOutputSize     string `json:"5. Output Size,omitempty"`
		SixTimeZone        string `json:"6. Time Zone,omitempty"`
	} `json:"Meta Data,omitempty"`
	TimeSeries5Min map[string]Ticker `json:"Time Series (5min),omitempty"`
}

func (c alphaVantage) IntraDay(ctx context.Context, symbol, month string) (*IntraDayResponse, error) {
	url := c.baseUrl

	q := url.Query()

	q.Add("function", "TIME_SERIES_INTRADAY")
	q.Add("symbol", symbol)
	q.Add("interval", "5min")
	q.Add("outputsize", "full")
	q.Add("month", month)
	q.Add("apikey", ENV_API_KEY)

	url.RawQuery = q.Encode()

	r, err := c.client.Request(ctx, http.MethodGet, url.String(), map[string]string{}, nil)
	if err != nil {
		return nil, err
	}

	res := &IntraDayResponse{}
	err = json.NewDecoder(r.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
