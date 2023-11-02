package client

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	env "github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rohanraj7316/ds/storage"
	"github.com/rohanraj7316/ds/utils/postgres"
)

const (
	envFile = ".env"
)

var loadEnv = env.Load

func TestInsertTicker(t *testing.T) {

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	st, err := NewStock()
	if err != nil {
		t.Errorf("failed to init stock api: %s", err)
	}

	res, err := st.IntraDay(context.Background(), "AAPL", "2023-10")
	if err != nil {
		t.Errorf("failed to get response from stock api: %s", err)
	}

	ps, err := postgres.New()
	if err != nil {
		t.Errorf("failed to init db: %s", err)
	}

	errs := []error{}
	tickers := []*storage.Ticker{}
	for key, val := range res.TimeSeries5Min {
		t, err := time.Parse(CONST_INTRA_DAY_DATE_FORMAT, key)
		if err != nil {
			errs = append(errs, err)
		}

		ticker := &storage.Ticker{
			Symbol:            "AAPL",
			Timestamp:         t.Unix(),
			ReadableTimestamp: t.Format(CONST_INTRA_DAY_DATE_FORMAT),
			OpenPrice:         val.OneOpen,
			HighPrice:         val.TwoHigh,
			LowPrice:          val.ThreeLow,
			ClosePrice:        val.FourClose,
			Volume:            val.FiveVolume,
		}

		tickers = append(tickers, ticker)
	}

	err = ps.Db.Create(tickers).Error
	if err != nil {
		t.Errorf("failed to insert record: %s", err)
	}
}

func ProcessDayNews(ctx context.Context, st Intelligence, ps *postgres.Storage,
	stock, timeFrom, timeTo string) error {

	res, err := st.NewsSentiments(context.Background(), stock, timeTo, timeFrom)
	if err != nil {
		errors.Wrapf(err, "failed to get response from stock api")
	}

	sentiments := []*storage.NewsSentiment{}
	adjustedTickers := []*storage.AdjustedTicker{}
	for _, val := range res.Feeds {
		tt, err := time.Parse(CONST_NEWS_SENTIMENTAL_DATE_FORMAT, val.TimePublished)
		if err != nil {
			errors.Wrapf(err, "failed to publish time into int")
		}

		from := tt.Add(time.Minute * -5).Format(CONST_INTRA_DAY_DATE_FORMAT)
		to := tt.Add(time.Minute * 5).Format(CONST_INTRA_DAY_DATE_FORMAT)

		tts := []storage.Ticker{}
		err = ps.Db.Raw("select id from tickers where readable_timestamp > ? AND readable_timestamp < ?", from, to).Find(&tts).Error
		if err != nil {
			errors.Wrapf(err, "unable to find records from: %s and to: %s", from, to)
		}

		fmt.Printf("tts: %+v\n\n", tts)

		sentiment := &storage.NewsSentiment{
			Symbol:                   stock,
			Timestamp:                tt.Unix(),
			ReadableTimestamp:        tt.Format(CONST_INTRA_DAY_DATE_FORMAT),
			SentimentScore:           val.OverallSentimentScore,
			AssociatedIntradayDataID: tts[0].ID,
		}

		sentiments = append(sentiments, sentiment)

		stt := tts[0]
		adjustedTicker := &storage.AdjustedTicker{
			Symbol:            stt.Symbol,
			Timestamp:         stt.Timestamp,
			ReadableTimestamp: stt.ReadableTimestamp,
			OpenPrice:         stt.OpenPrice,
			HighPrice:         stt.HighPrice,
			LowPrice:          stt.LowPrice,
			Volume:            stt.Volume,
			IntradayDataID:    stt.ID,
		}

		cp, _ := strconv.ParseFloat(stt.ClosePrice, 64)
		diff := float64(0.00)

		switch val.OverallSentimentLabel {
		case "Neutral":
		case "Somewhat-Bullish":
			diff = (cp * 0.01)
		case "Bullish":
			diff = (cp * 0.02)
		case "Somewhat-Bearish":
			diff = (cp * -0.01)
		case "Bearish":
			diff = (cp * -0.02)
		}

		adjustedTicker.ClosePrice = strconv.FormatFloat(cp+diff, 'f', -1, 64)

		adjustedTickers = append(adjustedTickers, adjustedTicker)
	}

	fmt.Printf("sentiments: %+v\n\n", sentiments)

	fmt.Printf("adjustedTickers: %+v\n\n", adjustedTickers)

	err = ps.Db.Create(sentiments).Error
	if err != nil {
		errors.Wrapf(err, "failed to insert record")
	}

	err = ps.Db.Create(adjustedTickers).Error
	if err != nil {
		errors.Wrapf(err, "failed to insert record")
	}

	return nil
}

func TestInsertAdjustedTicker(t *testing.T) {
	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	st, err := NewIntelligence()
	if err != nil {
		t.Errorf("failed to init stock api: %s", err)
	}

	ps, err := postgres.New()
	if err != nil {
		t.Errorf("failed to init db: %s", err)
	}

	// now := time.Now()

	// currentYear, currentMonth, _ := now.Date()

	// currentLocation := now.Location()

	// // Get the first day of the current month
	// firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)

	// // Get the last day of the previous month
	// lastOfMonth := firstOfMonth.AddDate(0, 0, -1)

	// // Get the first day of the previous month
	// firstOfLastMonth := lastOfMonth.AddDate(0, -1, 1)

	// start := firstOfLastMonth
	// for start.Before(lastOfMonth) {
	// 	start = start.AddDate(0, 0, 1)

	// 	from := start.Format(CONST_NEWS_SENTIMENTAL_DATE_FORMAT)
	// 	to := start.AddDate(0, 0, 1).Format(CONST_NEWS_SENTIMENTAL_DATE_FORMAT)

	err = ProcessDayNews(context.Background(), st, ps, "AAPL", "20231003T000000", "20231004T000000")
	if err != nil {
		t.Errorf("failed to update news data: %s", err)
	}

	t.Error()
	// }
}
