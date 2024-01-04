package voice_call

import (
	"time"

	"github.com/JoVi0li/ocean-server/internal/user"
	"github.com/google/uuid"
)

type VoiceCall struct {
	ID           uuid.UUID    `json:"-"`
	Participants [2]user.User `json:"participants"`
	StartedAt    time.Time    `json:"startedAt"`
	FinishedAt   time.Time    `json:"finishedAt"`
}
