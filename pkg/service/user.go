package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/utils.go"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
)

type UserService struct {
	repoUser  repository.User
	repoStore repository.Store
	repoMinio repository.FileServer
}

func NewUserService(repoUser repository.User, repoStore repository.Store, repoMinio repository.FileServer) *UserService {
	return &UserService{repoUser: repoUser, repoStore: repoStore, repoMinio: repoMinio}
}

func (s *UserService) GetCurrentUser(ctx context.Context, userId int) (model.User, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if errors.Is(err, static.ErrBannedUser) {
		return model.User{}, err
	}
	if err != nil {
		return model.User{}, fmt.Errorf("GetCurrentUser error: %v", err)
	}
	user.Sanitize()

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	if user.PasswordHash != "" {
		if !user.IsValid() {
			return user, static.ErrInvalidUser
		}
		user.Salt = utils.GenerateUuid()
		user.PasswordHash = utils.GeneratePasswordHash(user.PasswordHash, user.Salt)
		err := s.repoUser.UpdateUserPassword(ctx, user)
		if err != nil {
			return model.User{}, err
		}
	}
	correctUser, err := s.repoUser.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	correctUser.Sanitize()

	return correctUser, nil
}

func (s *UserService) CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error) {
	fileName := utils.GenerateUuid()
	bucketName := utils.GetBucketName(userId)
	err := s.repoMinio.CreateBucket(ctx, bucketName)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}

	buffer := make([]byte, size)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	link, err := s.repoMinio.UploadFile(ctx, bucketName, fileName, contentType, file, size)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}
	currentUser, err := s.repoUser.GetUserById(ctx, userId)
	if errors.Is(err, static.ErrBannedUser) {
		return fileName, err
	}
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}
	if currentUser.ImagePaths == nil {
		currentUser.ImagePaths = &[]string{link}
	} else {
		*currentUser.ImagePaths = append(*currentUser.ImagePaths, link)
	}
	_, err = s.repoUser.UpdateUser(ctx, currentUser)
	if err != nil {
		return link, fmt.Errorf("CreateFile error: %v", err)
	}

	return link, err
}

func (s *UserService) DeleteFile(ctx context.Context, userId int, link string) error {
	currentUser, err := s.repoUser.GetUserById(ctx, userId)
	if errors.Is(err, static.ErrBannedUser) {
		return err
	}
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}
	if currentUser.ImagePaths == nil {
		return static.ErrNoFiles
	}

	err = s.repoMinio.DeleteFile(ctx, utils.GetBucketName(userId), link)
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}
	*currentUser.ImagePaths = utils.Remove(*currentUser.ImagePaths, link)

	_, err = s.repoUser.UpdateUser(ctx, currentUser)
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}

	return err
}
