package voice_call

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Insert(ctx context.Context, voiceCall VoiceCall) (VoiceCall, error)
	FindById(ctx context.Context, id uuid.UUID) (VoiceCall, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, voiceCall VoiceCall) error
}

type RepositoryPostgres struct {
	Connection *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, voiceCall VoiceCall) (VoiceCall, error) {
	err := r.Connection.QueryRow(
		ctx,
		"INSERT INTO voice_calls (participants, startedAt) ($1, $2) RETURNING id, participants, startedAt",
		voiceCall.Participants,
		voiceCall.StartedAt,
	).Scan(&voiceCall.ID, &voiceCall.Participants, &voiceCall.StartedAt)

	if err != nil {
		return VoiceCall{}, err
	}

	return voiceCall, nil
}

func (r *RepositoryPostgres) FindById(ctx context.Context, id uuid.UUID) (VoiceCall, error) {
	var voiceCall = VoiceCall{ID: id}
	err := r.Connection.QueryRow(
		ctx,
		"SELECT participants, startedAt, finishedAt FROM voice_calls WHERE id = $1",
		id,
	).Scan(&voiceCall.Participants, &voiceCall.StartedAt, &voiceCall.FinishedAt)

	if err == pgx.ErrNoRows {
		return VoiceCall{}, ErrorVoiceCallNotFound
	}

	if err != nil {
		return VoiceCall{}, err
	}

	return voiceCall, nil
}

func (r *RepositoryPostgres) DeleteById(ctx context.Context, id uuid.UUID) error {
	tag, err := r.Connection.Exec(
		ctx,
		"DELETE FROM voice_calls WHERE id = $1",
		id,
	)

	if tag.RowsAffected() == 0 {
		return ErrorVoiceCallNotFound
	}

	return err
}


func (r *RepositoryPostgres) Update(ctx context.Context, voiceCall VoiceCall) error {
	_, err := r.Connection.Exec(
		ctx,
		`UPDATE "voice_call" SET "finishedAt" = COALESCE($1, "finishedAt") WHERE id = $2`,
		voiceCall.ID,
		voiceCall.FinishedAt,
	)

	return err
}