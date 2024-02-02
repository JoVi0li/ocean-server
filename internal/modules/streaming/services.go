package streaming

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func (s Service) StartCall(ctx context.Context, users [2]uuid.UUID) (Call, error) {
	return s.Repository.CreateCall(ctx, users)
}

func (s Service) FinishCall(ctx context.Context, id string) error {
	return s.Repository.UpdateCall(ctx, id)
}

func (s Service) GetCallById(ctx context.Context, id uuid.UUID) (Call, error) {
	return s.Repository.GetCallById(ctx, id)
}

