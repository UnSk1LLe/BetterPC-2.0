package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

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

func (s *AuthService) GenerateTokens(email, password string) {

}

func (s *AuthService) ParseToken(token string) {

}

func (s *AuthService) VerifyToken(token string) {

}
