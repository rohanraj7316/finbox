package migrations

import (
	"context"

	"github.com/rohanraj7316/ds/storage"
	"github.com/rohanraj7316/ds/utils/postgres"
)

func Migration(ctx context.Context, ps *postgres.Storage) error {
	return ps.Migrate(ctx,
		&storage.Ticker{},
		&storage.NewsSentiment{},
		&storage.AdjustedTicker{},
	)
}
