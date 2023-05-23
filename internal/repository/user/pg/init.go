package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	errorCommon "synapsis-challange/common/error"
	"synapsis-challange/common/sqlc"
	userModel "synapsis-challange/internal/model/user"
	userRepo "synapsis-challange/internal/repository/user"
)

type pgUserRepository struct {
	querier sqlc.Querier
}

func NewPGUserRepository(querier sqlc.Querier) userRepo.Repository {
	return &pgUserRepository{querier: querier}
}

func (r pgUserRepository) CreateUser(ctx context.Context, arg userModel.AddUser) (int32, error) {
	rid, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams(arg))
	if err != nil {
		return 0, err
	}
	return rid, err
}

func (r pgUserRepository) GetUser(ctx context.Context, id int32) (userModel.User, error) {
	user, err := r.querier.GetUser(ctx, id)
	if err == pgx.ErrNoRows {
		return userModel.User{}, errorCommon.NewNotFoundError("User not found")
	}
	return userModel.User(user), err
}

func (r pgUserRepository) GetUserByEmail(ctx context.Context, email string) (userModel.User, error) {
	user, err := r.querier.GetUserByEmail(ctx, email)
	if err == pgx.ErrNoRows {
		return userModel.User{}, errorCommon.NewNotFoundError("User not found")
	}
	return userModel.User(user), err
}
func (r pgUserRepository) VerifyAvailableUser(ctx context.Context, email string) (bool, error) {
	_, err := r.querier.GetUserByEmail(ctx, email)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
