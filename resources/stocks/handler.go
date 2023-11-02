package stocks

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	broadcast "github.com/rohanraj7316/ds/pkg/protos/v1/trading"
	"github.com/rohanraj7316/ds/storage"
	"github.com/rohanraj7316/ds/utils/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	broadcast.UnimplementedBroadcastServiceServer

	ps *postgres.Storage
}

func (s *Handler) RegisterGrpc(srv *grpc.Server) error {
	broadcast.RegisterBroadcastServiceServer(srv, s)
	return nil
}

func (s *Handler) RegisterHttp(ctx context.Context, mux *runtime.ServeMux,
	client *grpc.ClientConn) error {
	return broadcast.RegisterBroadcastServiceHandler(ctx, mux, client)
}

func New(ps *postgres.Storage) (h *Handler, err error) {
	h = &Handler{
		ps: ps,
	}

	return h, nil
}

func (h Handler) SelectTickerWithAdjustedTicker(symbol string) string {
	return fmt.Sprintf("SELECT t.id, symbol, COALESCE(at.closed_price, t.closed) AS closedPrice FROM ticker t LEFT JOIN adjusted_ticker at ON t.id = at.associated_intraday_data_id WHERE t.symbol = %s", symbol)
}

func (h *Handler) Read(ctx context.Context, in *broadcast.BroadcastServiceReadRequest) (*broadcast.BroadcastServiceReadResponse, error) {

	tickers := []storage.AdjustedTicker{}

	err := h.ps.Db.Raw(h.SelectTickerWithAdjustedTicker(in.Symbol)).Find(&tickers).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	parsedTickers := []*broadcast.Ticker{}

	for _, s := range tickers {
		t := &broadcast.Ticker{
			Symbol:     in.Symbol,
			ClosePrice: s.ClosePrice,
			Timestamp:  s.ReadableTimestamp,
		}

		parsedTickers = append(parsedTickers, t)
	}

	out := &broadcast.BroadcastServiceReadResponse{
		Tickers: parsedTickers,
	}

	return out, nil
}
