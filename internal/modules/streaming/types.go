package streaming

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Call struct {
	ID         uuid.UUID `json:"id"`
	StartedAt  uuid.UUID `json:"startedAt"`
	FinishedAt uuid.UUID `json:"finishedAt"`
}

type UsersInCall struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userID"`
	VoiceCallID uuid.UUID `json:"voiceCallID"`
}

type Connection struct {
	VoiceCallID uuid.UUID
	Clients     Clients
}

type Clients = map[*gin.Context]chan []byte
