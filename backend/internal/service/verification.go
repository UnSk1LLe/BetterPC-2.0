package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/service/helpers/passwordHasher"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	"BetterPC_2.0/pkg/data/models/users/requests/patch"
	"BetterPC_2.0/pkg/logging"
	"BetterPC_2.0/pkg/tokensGen"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const VerificationTokenByteLength = 32
const DefaultVerificationTokenTTL = 24 * time.Hour

type VerificationService struct {
	repo                repository.Verification
	notificationService *NotificationService
	logger              *logging.Logger
}

func NewVerificationService(repo repository.Verification, notificationService *NotificationService, logger *logging.Logger) *VerificationService {
	return &VerificationService{
		repo:                repo,
		notificationService: notificationService, //TODO manage notifications from verification service itself adn not from handlers
		logger:              logger,
	}
}

func (vs *VerificationService) SetNewToken(email string, tokenTTL time.Duration) (string, error) {
	token, err := tokensGen.GenerateNewUUID(email)
	if err != nil {
		return "", errors.New("error generating token: " + err.Error())
	}

	tokenExpirationTime := primitive.NewDateTimeFromTime(time.Now().Add(tokenTTL))

	err = vs.repo.SetTokenByEmail(email, token, tokenExpirationTime)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (vs *VerificationService) VerifyUser(token string) error {
	user, err := vs.repo.GetUserByVerificationToken(token)
	if err != nil {
		return err
	}
	if user.Verification.IsVerified {
		return userErrors.ErrUserAlreadyVerified
	}

	emptyToken := ""
	verified := true

	input := patch.UpdateUserVerificationDataRequest{
		VerificationToken: &emptyToken,
		IsVerified:        &verified,
	}

	err = vs.repo.UpdateVerificationDataById(user.ID, input)

	if err != nil {
		return err
	}

	return nil
}

func (vs *VerificationService) IsVerifiedUser(email string) (bool, error) {
	return vs.repo.IsVerifiedUser(email)
}

func (vs *VerificationService) GenerateRecoveryToken(email string) (string, error) {
	newToken, err := tokensGen.GenerateNewUUID(email)
	if err != nil {
		return "", errors.New("error generating token: " + err.Error())
	}

	return newToken, nil
}

func (vs *VerificationService) UpdatePasswordByToken(token, password string) error {
	user, err := vs.repo.GetUserByVerificationToken(token)
	if err != nil {
		return err
	}

	if !user.Verification.IsVerified {
		return userErrors.ErrUserNotVerified
	}

	passwordHash, err := passwordHasher.GeneratePasswordHash(password)
	if err != nil {
		return err
	}

	err = vs.repo.UpdateUserPasswordById(user.ID, passwordHash)
	if err != nil {
		return err
	}

	return nil
}
