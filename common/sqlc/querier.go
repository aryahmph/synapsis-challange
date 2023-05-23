// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetUser(ctx context.Context, id int32) (GetUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
}

var _ Querier = (*Queries)(nil)
