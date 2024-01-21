package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Insert(ctx context.Context, user User) (User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type RepositoryPostgres struct {
	Connection *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, user User) (User, error) {
	err := r.Connection.QueryRow(
		ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	var user = User{ID: id}
	err := r.Connection.QueryRow(
		ctx,
		"SELECT username, email FROM users WHERE id = $1",
		id,
	).Scan(&user.Username, &user.Email)

	if err == pgx.ErrNoRows {
		return User{}, ErrorUserNotFound
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) FindByEmail(ctx context.Context, email string) (User, error) {
	var user = User{Email: email}
	err := r.Connection.QueryRow(
		ctx,
		"SELECT username, email, password, id FROM users WHERE email = $1",
		email,
	).Scan(&user.Username, &user.Email, &user.Password, &user.ID)

	if err == pgx.ErrNoRows {
		return User{}, ErrorUserNotFound
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) DeleteById(ctx context.Context, id uuid.UUID) error {
	tag, err := r.Connection.Exec(
		ctx,
		"DELETE FROM users WHERE id = $1",
		id,
	)

	if tag.RowsAffected() == 0 {
		return ErrorUserNotFound
	}

	return err
}
