package voice_call

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(ctx context.Context, voiceCall VoiceCall) (VoiceCall, error) {
	if len(voiceCall.Participants) != 2 {
		return VoiceCall{}, ErrorMissingParticipant
	}

	return s.Create(ctx, voiceCall)
}

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (VoiceCall, error) {
	return s.Repository.FindById(ctx, id)
}

func (s *Service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.Repository.DeleteById(ctx, id)
}

func (s *Service) Finish(ctx context.Context, id uuid.UUID) error {
	voiceCall, err := s.FindById(ctx, id)

	if err != nil {
		return err
	}

	if !voiceCall.FinishedAt.IsZero() {
		return ErrorVoiceCallAlreadyFinished
	}

	voiceCall.FinishedAt = time.Now()

	return s.Repository.Update(ctx, voiceCall)
}