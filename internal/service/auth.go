package service

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/users"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user users.User) (primitive.ObjectID, error) {
	user.UserInfo.Password = s.generatePasswordHash(user.UserInfo.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash)
}

func (s *AuthService) GenerateTokenPair(email, password string) (TokenPair, error) {
	user, err := s.repo.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return TokenPair{}, err
	}

	accessTokenString, err := generateAccessToken(&user)
	refreshTokenString, err := generateRefreshToken(&user)

	tokens := TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokens, nil
}

func generateAccessToken(user *users.User) (string, error) {
	//creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.GetConfig().Tokens.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.ID.Hex(),
		Email:  user.UserInfo.Email,
		Role:   user.UserInfo.Role,
	})

	accessTokenString, err := accessToken.SignedString([]byte(configs.GetConfig().Tokens.AccessTokenSigningKey))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error generating access token: %s", err))
	}

	return accessTokenString, nil
}

func generateRefreshToken(user *users.User) (string, error) {
	//creating refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.GetConfig().Tokens.RefreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.ID.Hex(),
		Email:  user.UserInfo.Email,
		Role:   user.UserInfo.Role,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(configs.GetConfig().Tokens.RefreshTokenSigningKey))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error generating refresh token: %s", err))
	}

	return refreshTokenString, nil
}

func (s *AuthService) ParseAccessToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return configs.GetConfig().Tokens.AccessTokenSigningKey, nil
	})
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return primitive.NilObjectID.Hex(), errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) RefreshTokens(refreshTokenString string) (TokenPair, error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.GetConfig().Tokens.RefreshTokenSigningKey, nil
	})
	if err != nil || !refreshToken.Valid {
		return TokenPair{}, errors.New("invalid or expired refresh token")
	}

	claims, ok := refreshToken.Claims.(*tokenClaims)
	if !ok || claims.UserId == "" {
		return TokenPair{}, errors.New("invalid token claims")
	}

	userID, err := primitive.ObjectIDFromHex(claims.UserId)
	if err != nil {
		return TokenPair{}, errors.New("invalid token claims: could not get Object ID from hex")
	}

	newAccessToken, err := generateAccessToken(&users.User{
		ID: userID,
		UserInfo: users.UserInfo{
			Email: claims.Email,
			Role:  claims.Role,
		},
	})
	if err != nil {
		return TokenPair{}, err
	}

	newRefreshToken, err := generateRefreshToken(&users.User{
		ID: userID,
		UserInfo: users.UserInfo{
			Email: claims.Email,
			Role:  claims.Role,
		},
	})
	if err != nil {
		return TokenPair{}, err
	}

	tokens := TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return tokens, nil
}
