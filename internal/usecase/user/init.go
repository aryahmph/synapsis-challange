package user

import (
	passwordCommon "synapsis-challange/common/password"
	userRepo "synapsis-challange/internal/repository/user"
)

type userUsecase struct {
	userRepository  userRepo.Repository
	passwordManager *passwordCommon.PasswordHashManager
}

func NewUserUsecase(
	uRepository userRepo.Repository,
	passwordManager *passwordCommon.PasswordHashManager,
) *userUsecase {
	return &userUsecase{
		userRepository:  uRepository,
		passwordManager: passwordManager,
	}
}
