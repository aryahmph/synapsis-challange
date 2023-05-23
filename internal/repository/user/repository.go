package user

import (
	"context"
	userModel "synapsis-challange/internal/model/user"
)

type Repository interface {
	CreateUser(ctx context.Context, arg userModel.AddUser) (int32, error)
	GetUser(ctx context.Context, id int32) (userModel.User, error)
	GetUserByEmail(ctx context.Context, email string) (userModel.User, error)
	VerifyAvailableUser(ctx context.Context, email string) (bool, error)
}
