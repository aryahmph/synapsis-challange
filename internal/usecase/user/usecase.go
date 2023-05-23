package user

import (
	"context"
	userModel "synapsis-challange/internal/model/user"
)

type Usecase interface {
	CreateUser(ctx context.Context, user userModel.AddUser) (int32, error)
	Login(ctx context.Context, email, password string) (string, error)
}
