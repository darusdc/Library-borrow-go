package service

import (
	"context"
	"errors"
	"time"

	"github.com/darusdc/belajar-go/config"
	"github.com/darusdc/belajar-go/domain"
	"github.com/darusdc/belajar-go/dto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

// Login implements domain.AuthService.
func (a authService) Login(ctx context.Context, usr dto.AuthRequest) (response dto.AuthResponse, err error) {
	user, err := a.userRepository.FindByEmail(ctx, usr.Email)

	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("authentication is failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("password is wrong")
	}

	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, errors.New("password is wrong")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}

func NewAuth(cnf *config.Config, usr domain.UserRepository) domain.AuthService {
	return authService{
		conf:           cnf,
		userRepository: usr,
	}
}
