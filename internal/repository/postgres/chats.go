package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type ChatRepository struct {
	*PostgresRepository
}

func (r ChatRepository) GetByIDs(ctx context.Context, ids []uuid.UUID) (*entity.DataResult, error) {
	var (
		channels []entity.Channel
		diologs  []entity.Diolog
		groups   []entity.Group
	)

	err := r.WithTx(ctx, func(tx pgx.Tx) error {
		var err error

		channels, err = GetChannelsTx(ctx, tx, ids)
		if err != nil {
			return fmt.Errorf("failed to get channels: %w", err)
		}

		diologs, err = GetDiologsTx(ctx, tx, ids)
		if err != nil {
			return fmt.Errorf("failed to get dialogs: %w", err)
		}

		groups, err = GetGroupsTx(ctx, tx, ids)
		if err != nil {
			return fmt.Errorf("failed to get groups: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("transaction error: %w", err)
	}

	return &entity.DataResult{
		Channels: channels,
		Diologs:  diologs,
		Groups:   groups,
	}, nil
}
