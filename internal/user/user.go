package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
