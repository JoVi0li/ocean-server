package streaming

import "context"

type Service struct {
	Repository Repository
}

func (s Service) StartCall(ctx context.Context, users [2]UsersInCall) (Call, error) {
	return s.Repository.CreateCall(ctx, users)
}

func (s Service) FinishCall(ctx context.Context, id string) error {
	return s.Repository.UpdateCall(ctx, id)
}
