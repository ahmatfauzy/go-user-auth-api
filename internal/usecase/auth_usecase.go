package usecase

import (
	"auth-api/internal/domain"
	"auth-api/pkg/jwt"
	"auth-api/pkg/redisclient"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo       domain.UserRepository
	jwtSecret      string
	jwtExpireHours int
	redis          *redisclient.RedisClient
}

func NewAuthUsecase(repo domain.UserRepository, secret string, expire int, redis *redisclient.RedisClient) domain.AuthUseCase {
	return &authUsecase{
		userRepo:       repo,
		jwtSecret:      secret,
		jwtExpireHours: expire,
		redis:          redis,
	}
}

func (uc *authUsecase) Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{Email: email, Password: string(hashed)}
	return uc.userRepo.Create(user)
}

func (uc *authUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}
	return jwt.GenerateToken(user.ID, uc.jwtSecret, time.Duration(uc.jwtExpireHours)*time.Hour)
}

func (uc *authUsecase) RefreshToken(token string) (string, error) {
	claims, err := jwt.ValidateToken(token, uc.jwtSecret)
	if err != nil {
		return "", err
	}
	return jwt.GenerateToken(claims.UserID, uc.jwtSecret, time.Duration(uc.jwtExpireHours)*time.Hour)
}

func (uc *authUsecase) Logout(token string) error {
	claims, err := jwt.ValidateToken(token, uc.jwtSecret)
	if err != nil {
		return err
	}
	exp := time.Until(time.Unix(claims.ExpiresAt.Unix(), 0))
	return uc.redis.Set("blacklist:"+token, "1", exp)
}