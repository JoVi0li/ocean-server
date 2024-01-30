package streaming

import (
	"context"

	"github.com/JoVi0li/ocean-server/internal/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateCall(ctx context.Context, call Call, users []UsersInCall) error
	UpdateCall(ctx context.Context, call Call) error
}

type RepositoryPostgres struct {
	Connection *pgxpool.Pool
}

func (r *RepositoryPostgres) CreateCall(ctx context.Context, call Call, users [2]UsersInCall) (Call, error) {
	err := r.Connection.QueryRow(
		ctx,
		"BEGIN;INSERT INTO voice_calls RETURNING id INTO result_id, started_at, finished_at;ROLLBACK;INSERT INTO users_in_calls (user_id, voice_call_id) VALUES ($1, result_id);ROLLBACK;INSERT INTO users_in_calls (user_id, voice_call_id) VALUES ($2, result_id);ROLLBACK;COMMIT;",
		users[0].ID,
		users[1].ID,
	).Scan(&call.ID, &call.StartedAt, &call.FinishedAt)
	if err != nil {
		return Call{}, err
	}

	return call, nil
}

func (r *RepositoryPostgres) UpdateCall(ctx context.Context, call Call) error {
	tag, err := r.Connection.Exec(
		ctx,
		"UPDATE voice_calls SET finished_at = CURRENT_TIMESTAMP WHERE id = $1 AND finished_at IS NULL;",
		call.ID,
	)
	if tag.RowsAffected() == 0 {
		return shared.ErrorUserNotFound
	}

	return err
}
