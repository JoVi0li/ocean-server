package streaming

import "github.com/google/uuid"

type VoiceCall struct {
	ID         uuid.UUID `json:"id"`
	StartedAt  uuid.UUID `json:"startedAt"`
	FinishedAt uuid.UUID `json:"finishedAt"`
}

type UsersInCall struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userID"`
	VoiceCallID uuid.UUID `json:"voiceCallID"`
}
