package repository

import (
	"context"

	"github.com/google/uuid"

	"CRUD-GO/db"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		q: q,
	}
}

func (ur *UserRepository) UpdateUserName(ctx context.Context, newUser *db.User) error {
	return ur.q.UpdateUsersName(ctx, db.UpdateUsersNameParams{
		ID:   newUser.ID,
		Name: newUser.Name,
	})
}

func (ur *UserRepository) CreateUser(ctx context.Context, name string) error {
	return ur.q.CreateUser(ctx, db.CreateUserParams{
		ID:   uuid.New(),
		Name: name,
	})
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ur.q.DeleteUser(ctx, id)
}

func (ur *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (db.User, error) {
	return ur.q.GetUserById(ctx, id)
}

func (ur *UserRepository) ListAllUsers(ctx context.Context) ([]db.User, error) {
	return ur.q.GetUsers(ctx)
}
