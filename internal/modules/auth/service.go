package auth

import (
	"context"
	"net/mail"

	"github.com/JoVi0li/ocean-server/internal/shared"
	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func (s Service) Create(ctx context.Context, user User) (User, error) {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return User{}, shared.ErrorEmailInvalid
	}

	invalidUsername, invalidPassword := len(user.Username) < 3, len(user.Password) < 8
	if invalidUsername {
		return User{}, shared.ErrorUsernameInvalid
	}
	if invalidPassword {
		return User{}, shared.ErrorPasswordInvalid
	}

	hashedPass, err := shared.HashPassword(user.Password)
	if err != nil {
		return User{}, shared.ErrorTryingHashPassword
	}

	user.Password = hashedPass

	return s.Repository.CreateUser(ctx, user)
}

func (s Service) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	return s.Repository.GetUserById(ctx, id)
}

func (s Service) FindByEmail(ctx context.Context, email string) (User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return User{}, shared.ErrorEmailInvalid
	}

	return s.Repository.GetUserByEmail(ctx, email)
}

func (s Service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.Repository.DeleteUserById(ctx, id)
}
