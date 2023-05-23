package user

import (
	"context"
	userModel "synapsis-challange/internal/model/user"
)

type Usecase interface {
	CreateUser(ctx context.Context, user userModel.AddUser) (int32, error)
	LoginUser(ctx context.Context, email, password string) (string, error)
}
