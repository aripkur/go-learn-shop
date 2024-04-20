package auth

import (
	"context"

	"github.com/aripkur/go-learn-shop/infra/response"
	"github.com/aripkur/go-learn-shop/internal/config"
)

type Repository interface {
	GetAuthByEmail(ctx context.Context, email string) (AuthEntity, error)
	CreateAuth(ctx context.Context, model AuthEntity) error
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
	authReq := newFromRegisterRequest(req)
	if err = authReq.validate(); err != nil {
		return
	}

	if err = authReq.encryptPassword(config.Cfg.App.Encryption.Salt); err != nil {
		return
	}

	authEntity, err := s.repo.GetAuthByEmail(ctx, authReq.Email)

	if err != nil && err != response.ErrNotFound {
		return
	}

	if authEntity.isExists() {
		return response.ErrEmailAlReadyUsed
	}

	return s.repo.CreateAuth(ctx, authReq)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authReq := newFromLoginRequest(req)
	if err = authReq.validateEmail(); err != nil {
		return
	}

	if err = authReq.validatePassword(); err != nil {
		return
	}

	authEntity, err := s.repo.GetAuthByEmail(ctx, authReq.Email)
	if err != nil {
		return
	}

	// plaintext in authReq
	if err = authReq.VerifyPasswordFromPlain(authEntity.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = authEntity.GenerateToken(config.Cfg.App.Encryption.JwtSecret)

	return
}
