package user

import (
	"context"
	"net/mail"

	"github.com/JoVi0li/ocean-server/internal/util"
	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func (s Service) Create(ctx context.Context, user User) (User, error) {
	_, errInvalidEmail := mail.ParseAddress(user.Email)
	
	if errInvalidEmail != nil {
		return User{}, ErrorEmailInvalid
	}

	if len(user.Username) < 3 {
		return User{}, ErrorUsernameInvalid
	}

	if len(user.Password) < 8 {
		return User{}, ErrorPasswordInvalid
	}

	hashedPass, errHashPass := util.HashPassword(user.Password)

	if errHashPass != nil {
		return User{}, ErrorTryingHashPassword
	}

	user.Password = hashedPass

	return s.Repository.Insert(ctx, user)
}

func (s Service) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	return s.Repository.FindById(ctx, id)
}

func (s Service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.Repository.DeleteById(ctx, id)
}
