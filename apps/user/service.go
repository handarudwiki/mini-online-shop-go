package user

import (
	"context"
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/internal/config"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (user UserEntity, err error)
	CreateUser(ctx context.Context, user UserEntity) (err error)
}
type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	userEntity := NewFormRegisterRequest(req)

	err = userEntity.Validate()

	if err != nil {
		return
	}

	user, err := s.repo.GetUserByEmail(ctx, userEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if user.IsExists() {
		return response.ErrEmailAlreadyUsed
	}

	err = userEntity.HashPassword(config.Cfg.App.Encryption.Salt)

	if err != nil {
		return
	}

	err = s.repo.CreateUser(ctx, userEntity)
	if err != nil {
		return
	}
	return
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	userEntity := NewFormLoginRequest(req)

	err = userEntity.Validate()

	if err != nil {
		return
	}

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return
	}

	err = user.VerifyPassword(req.Password)
	if err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = user.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	if err != nil {
		return
	}

	return

}
