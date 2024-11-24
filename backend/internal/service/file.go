package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileServiceError struct {
	Operation string
	Err       error
}

func (e FileServiceError) Error() string {
	return fmt.Sprintf("FileServiceError during %s: %s", e.Operation, e.Err.Error())
}

func (e FileServiceError) Unwrap() error {
	return e.Err
}

var (
	ErrUnsupportedFormat  = errors.New("unsupported file format")
	ErrCouldNotCreateFile = errors.New("error creating file")
	ErrCouldNotDeleteFile = errors.New("error deleting file")
)

const (
	imgDirectoryName        = "images"
	userImgDirectoryName    = "users"
	productImgDirectoryName = "products"
)

type FileService struct {
	StaticPath string
}

func NewFileService(staticPath string) *FileService {
	return &FileService{StaticPath: staticPath}
}

func (fileService *FileService) AddUserImage(file *multipart.FileHeader) (string, error) {
	return fileService.AddImage(file, userImgDirectoryName)
}

func (fileService *FileService) DeleteUserImage(imageName string) error {
	return fileService.DeleteImage(imageName, userImgDirectoryName)
}

func (fileService *FileService) AddProductImage(file *multipart.FileHeader) (string, error) {
	return fileService.AddImage(file, productImgDirectoryName)
}

func (fileService *FileService) DeleteProductImage(imageName string) error {
	return fileService.DeleteImage(imageName, productImgDirectoryName)
}

func (fileService *FileService) AddImage(file *multipart.FileHeader, subDirectory string) (string, error) {
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".bmp":  true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", FileServiceError{Operation: "AddImage", Err: errors.Wrap(ErrUnsupportedFormat, "allowed formats are: .png, .jpg, .jpeg, .bmp")}
	}

	imgPath := filepath.Join(fileService.StaticPath, imgDirectoryName, subDirectory)
	if err := ensureDirectoryExists(imgPath); err != nil {
		return "", FileServiceError{Operation: "AddImage", Err: errors.Wrap(ErrCouldNotCreateFile, err.Error())}
	}

	fileName := uuid.New().String() + ext
	destination := filepath.Join(imgPath, fileName)

	if err := addFile(file, destination); err != nil {
		return "", FileServiceError{Operation: "AddImage", Err: errors.Wrapf(ErrCouldNotCreateFile, "failed to add image from file: %s", file.Filename)}
	}

	return fileName, nil
}

func (fileService *FileService) DeleteImage(imageName, subDirectory string) error {
	imagePath := filepath.Join(fileService.StaticPath, imgDirectoryName, subDirectory, imageName)

	err := deleteFile(imagePath)
	if err != nil {
		return FileServiceError{Operation: "DeleteImage", Err: errors.Wrapf(err, "failed to delete image: %s", imageName)}
	}

	return nil
}

func ensureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func addFile(file *multipart.FileHeader, destination string) error {
	srcFile, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func deleteFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
