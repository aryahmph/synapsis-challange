package user

import (
	jwtCommon "synapsis-challange/common/jwt"
	passwordCommon "synapsis-challange/common/password"
	userRepo "synapsis-challange/internal/repository/user"
)

type userUsecase struct {
	userRepository  userRepo.Repository
	jwtManager      *jwtCommon.JWTManager
	passwordManager *passwordCommon.PasswordHashManager
}

func NewUserUsecase(
	uRepository userRepo.Repository,
	jwtManager *jwtCommon.JWTManager,
	passwordManager *passwordCommon.PasswordHashManager,
) *userUsecase {
	return &userUsecase{
		userRepository:  uRepository,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}
