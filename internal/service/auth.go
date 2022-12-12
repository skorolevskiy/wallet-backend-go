package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
	"math/rand"
	"strconv"
	"time"
)

const (
	salt           = "^3cH9s72^aM@Z-fM"
	signingKey     = "1^3cH9s72^aM@Z-fM"
	tokenTTL       = 15 * time.Minute
	sessionExpires = time.Hour * 24 * 30
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

type AuthService struct {
	repo      repository.Authorization
	repoToken repository.Tokens
}

func NewAuthService(repo repository.Authorization, repoToken repository.Tokens) *AuthService {
	return &AuthService{
		repo:      repo,
		repoToken: repoToken,
	}
}

func (s *AuthService) CreateUser(user domain.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) SignIn(username, password string) (string, string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", "", err
	}

	return s.generateTokens(user.ID)
}

func (s *AuthService) generateTokens(userId int64) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			Subject:   strconv.Itoa(int(userId)),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		userId,
	})
	accessToken, err := t.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	if err := s.repoToken.Create(domain.RefreshToken{
		UserID:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(sessionExpires),
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	session, err := s.repoToken.Get(refreshToken)
	if err != nil {
		return "", "", err
	}

	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", errors.New("token expired")
	}

	return s.generateTokens(session.UserID)
}
