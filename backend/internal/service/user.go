package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/service/helpers/converters"
	"BetterPC_2.0/internal/service/helpers/passwordHasher"
	"BetterPC_2.0/pkg/data/models/users"
	userFilters "BetterPC_2.0/pkg/data/models/users/filters"
	"BetterPC_2.0/pkg/data/models/users/requests/admin"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"BetterPC_2.0/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"time"
)

type UserService struct {
	repo        repository.User
	fileService *FileService
	logger      *logging.Logger
}

func NewUserService(repo repository.User, fileService *FileService, logger *logging.Logger) *UserService {
	return &UserService{
		repo:        repo,
		fileService: fileService,
		logger:      logger,
	}
}

func (u *UserService) Create(input adminUserRequest.CreateUserRequest) (string, error) {

	dob, err := converters.ConvertDateFromString(input.Dob)
	if err != nil {
		return "", err
	}

	passwordHash, err := passwordHasher.GeneratePasswordHash(input.Password)
	if err != nil {
		return "", err
	}

	currentTime := primitive.NewDateTimeFromTime(time.Now())

	user := users.User{
		UserInfo: users.UserInfo{
			Name:     input.Name,
			Surname:  input.Surname,
			Dob:      primitive.NewDateTimeFromTime(dob),
			Email:    input.Email,
			Password: passwordHash,
			Role:     input.Role,
			Image:    "",
		},
		Verification: users.Verification{
			Token:      "",
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
			ExpiresAt:  currentTime,
			IsVerified: true,
		},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	userObjId, err := u.repo.Create(user)
	if err != nil {
		return "", err
	}

	return userObjId.String(), nil
}

func (u *UserService) SetRole(userId string, role string) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	newRole, err := users.UserRoleFromString(role)
	if err != nil {
		return err
	}

	return u.repo.SetRole(userObjId, newRole)
}

func (u *UserService) UpdateUserInfo(userId string, input userUpdateRequests.UpdateUserInfoRequest) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	return u.repo.UpdateUserInfoById(userObjId, input)
}

func (u *UserService) UpdateUserImage(userId string, image *multipart.FileHeader) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	imgUrl, err := u.fileService.AddUserImage(image)
	if err != nil {
		return err
	}

	prevImgUrl, err := u.repo.UpdateUserImageById(userObjId, imgUrl)
	if err != nil {
		return err
	}

	err = u.fileService.DeleteUserImage(prevImgUrl)
	if err != nil {
		u.logger.Error(err.Error())
	}

	return nil
}

func (u *UserService) Delete(userId string) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	return u.repo.DeleteById(userObjId)
}

func (u *UserService) GetList(filters userFilters.AdminUserFilters) ([]users.User, error) {
	var bsonFilter bson.M

	if filters.DateFrom != nil && filters.DateTo != nil {
		bsonFilter["created_at"] = bson.M{
			"$gte": *filters.DateFrom,
			"$lte": *filters.DateTo,
		}
	} else if filters.DateFrom != nil {
		bsonFilter["created_at"] = bson.M{"$gte": *filters.DateFrom}
	} else if filters.DateTo != nil {
		bsonFilter["created_at"] = bson.M{"$lte": *filters.DateTo}
	}

	if filters.IsVerified != nil {
		bsonFilter["verification.is_verified"] = *filters.IsVerified
	}

	if len(filters.Roles) > 0 {
		bsonFilter["user_info.roles"] = bson.M{"$in": filters.Roles}
	}

	return u.repo.GetList(bsonFilter)
}

func (u *UserService) GetById(userId string) (users.User, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return users.User{}, err
	}

	return u.repo.GetById(userObjId)
}
