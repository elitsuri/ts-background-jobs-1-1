package service

import (
	"context"
	"errors"

	"github.com/example/ts-background-jobs-1/internal/domain"
	"github.com/example/ts-background-jobs-1/internal/repository"
	"github.com/example/ts-background-jobs-1/pkg/hash"
	"github.com/example/ts-background-jobs-1/pkg/jwt"
)

var ErrEmailTaken = errors.New("email already taken")
var ErrInvalidCredentials = errors.New("invalid credentials")

// AuthService encapsulates authentication business logic.
type AuthService struct {
	users *repository.UserRepository
	token *jwt.Manager
}

func NewAuthService(users *repository.UserRepository, token *jwt.Manager) *AuthService {
	return &AuthService{users: users, token: token}
}

func (s *AuthService) Register(ctx context.Context, req domain.RegisterRequest) (*domain.UserPublic, error) {
	exists, err := s.users.ExistsByEmail(ctx, req.Email)
	if err != nil { return nil, err }
	if exists { return nil, ErrEmailTaken }
	hashed, err := hash.Password(req.Password)
	if err != nil { return nil, err }
	user, err := s.users.Create(ctx, req.Email, req.FullName, hashed)
	if err != nil { return nil, err }
	return &domain.UserPublic{ID: user.ID, Email: user.Email, FullName: user.FullName, CreatedAt: user.CreatedAt}, nil
}

func (s *AuthService) Login(ctx context.Context, req domain.LoginRequest) (*domain.TokenPair, error) {
	user, hashed, err := s.users.FindByEmail(ctx, req.Email)
	if err != nil { return nil, err }
	if user == nil || !hash.Verify(req.Password, hashed) {
		return nil, ErrInvalidCredentials
	}
	access, err := s.token.Create(user.ID)
	if err != nil { return nil, err }
	return &domain.TokenPair{AccessToken: access, TokenType: "bearer"}, nil
}
