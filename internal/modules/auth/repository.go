package auth

import (
	"context"

	"github.com/JoVi0li/ocean-server/internal/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	DeleteUserById(ctx context.Context, id uuid.UUID) error
}

type RepositoryPostgres struct {
	Connection *pgxpool.Pool
}

func (r *RepositoryPostgres) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.Connection.QueryRow(
		ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email;",
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	var user = User{ID: id}

	err := r.Connection.QueryRow(
		ctx,
		"SELECT id, username, email, password FROM users WHERE id = $1;",
		id,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == pgx.ErrNoRows {
		return User{}, shared.ErrorUserNotFound
	}
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user = User{Email: email}

	err := r.Connection.QueryRow(
		ctx,
		"SELECT id, username, email, password FROM users WHERE email = $1;",
		email,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == pgx.ErrNoRows {
		return User{}, shared.ErrorUserNotFound
	}
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	tag, err := r.Connection.Exec(
		ctx,
		"DELETE FROM users WHERE id = $1;",
		id,
	)
	if tag.RowsAffected() == 0 {
		return shared.ErrorUserNotFound
	}

	return err
}
