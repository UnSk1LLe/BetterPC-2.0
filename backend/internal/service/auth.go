package service

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/service/helpers/passwordHasher"
	"BetterPC_2.0/pkg/data/models/users"
	userRequests "BetterPC_2.0/pkg/data/models/users/requests/auth"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/tokens"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	accessSigningKey  string
	accessTTL         time.Duration
	refreshSigningKey string
	refreshTTL        time.Duration
)

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type tokenInput struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type tokenClaims struct {
	jwt.RegisteredClaims
	tokenInput
}

type AuthService struct {
	repo     repository.Authorization
	userRepo repository.User
}

func InitAuth(cfg *configs.Config) {
	accessSigningKey = cfg.Tokens.AccessTokenSigningKey
	accessTTL = cfg.Tokens.AccessTokenTTL
	refreshSigningKey = cfg.Tokens.RefreshTokenSigningKey
	refreshTTL = cfg.Tokens.RefreshTokenTTL
}

func NewAuthService(repo repository.Authorization, userRepo repository.User) *AuthService {
	return &AuthService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *AuthService) CreateUser(input userRequests.RegisterRequest) (primitive.ObjectID, error) {

	dob, err := time.Parse("2006-01-02", input.Dob) //Parse dob string to time.Time
	if err != nil {
		return primitive.NilObjectID, errors.New("invalid date of birth format")
	}

	passwordHash, err := passwordHasher.GeneratePasswordHash(input.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}

	verificationToken, err := tokens.GenerateNewUUID(input.Email)
	if err != nil {
		return primitive.NilObjectID, errors.New("failed to generate verification token: " + err.Error())
	}

	user := users.NewUserDefault(verificationToken, configs.GetConfig()) //TODO generate verificaitonToken with verification service
	user.UserInfo.Email = input.Email
	user.UserInfo.Name = input.Name
	user.UserInfo.Surname = input.Surname
	user.UserInfo.Dob = primitive.NewDateTimeFromTime(dob)
	user.UserInfo.Password = passwordHash

	return s.repo.CreateUser(*user)
}

func (s *AuthService) GenerateTokenPair(email, password string) (TokenPair, error) {

	user, err := s.repo.GetUserByEmail(email)

	if err != nil {
		return TokenPair{}, err
	}

	// Compare the provided password with the stored hash
	err = passwordHasher.CompareHashAndPassword(user.UserInfo.Password, password)

	if err != nil {
		return TokenPair{}, errors.New("invalid credentials")
	}

	input := tokenInput{
		UserId: user.ID.Hex(),
		Email:  user.UserInfo.Email,
		Role:   user.UserInfo.Role,
	}

	accessTokenString, err := s.generateAccessToken(&input)
	refreshTokenString, err := s.generateRefreshToken(&input)

	tokens := TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokens, nil
}

func (s *AuthService) ParseAccessToken(accessToken string) (userResponses.UserResponse, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(accessSigningKey), nil
	})
	if err != nil {
		return userResponses.UserResponse{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return userResponses.UserResponse{}, errors.New("token claims are not of type *tokenClaims")
	}

	userObjId, err := primitive.ObjectIDFromHex(claims.UserId)
	if err != nil {
		return userResponses.UserResponse{}, err
	}

	user, err := s.userRepo.GetById(userObjId)
	if err != nil {
		return userResponses.UserResponse{}, err
	}

	return user.ConvertToUserResponse(), nil
}

func (s *AuthService) generateAccessToken(input *tokenInput) (string, error) {
	//creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		tokenInput: tokenInput{
			UserId: input.UserId,
			Email:  input.Email,
			Role:   input.Role,
		},
	})

	accessTokenString, err := accessToken.SignedString([]byte(accessSigningKey))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error generating access token: %s", err))
	}

	return accessTokenString, nil
}

func (s *AuthService) generateRefreshToken(input *tokenInput) (string, error) {
	//creating refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		tokenInput: tokenInput{
			UserId: input.UserId,
			Email:  input.Email,
			Role:   input.Role,
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshSigningKey))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error generating refresh token: %s", err))
	}

	return refreshTokenString, nil
}

func (s *AuthService) RefreshTokens(refreshTokenString string) (userResponses.UserResponse, TokenPair, error) {
	var response userResponses.UserResponse

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(refreshSigningKey), nil
	})

	if err != nil || !refreshToken.Valid {
		return response, TokenPair{}, errors.New("invalid or expired refresh token")
	}

	claims, ok := refreshToken.Claims.(*tokenClaims)
	if !ok || claims.UserId == "" {
		return response, TokenPair{}, errors.New("invalid token claims")
	}

	userID, err := primitive.ObjectIDFromHex(claims.UserId)
	if err != nil {
		return response, TokenPair{}, errors.New("invalid token claims: could not get Object ID from hex")
	}

	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return response, TokenPair{}, err
	}

	newAccessToken, err := s.generateAccessToken(&tokenInput{
		UserId: claims.UserId,
		Email:  claims.Email,
		Role:   claims.Role,
	})
	if err != nil {
		return response, TokenPair{}, err
	}

	newRefreshToken, err := s.generateRefreshToken(&tokenInput{
		UserId: claims.UserId,
		Email:  claims.Email,
		Role:   claims.Role,
	})
	if err != nil {
		return response, TokenPair{}, err
	}

	tokens := TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	response = user.ConvertToUserResponse()
	return response, tokens, nil
}

func (s *AuthService) HasRole(userId primitive.ObjectID, roles ...string) (bool, error) {
	if len(roles) == 0 {
		return false, errors.New("no roles provided for check provided: argument <roles> must contain at least one value")
	}

	hasRole, err := s.repo.HasRole(userId, roles)
	if err != nil {
		return false, err
	}
	return hasRole, nil
}
