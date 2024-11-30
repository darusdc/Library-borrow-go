package domain

import (
	"context"

	"github.com/darusdc/belajar-go/dto"
)

type AuthService interface {
	Login(ctx context.Context, usr dto.AuthRequest) (dto.AuthResponse, error)
}
