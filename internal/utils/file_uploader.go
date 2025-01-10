package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// UploadFile uploads a file and renames it with the movie ID
func UploadFile(file *multipart.FileHeader, movieID uuid.UUID) (string, error) {
	// Create the directory if it doesn't exist
	dir := "./assets/file/movie/"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// Create the file path
	fileName := fmt.Sprintf("%s%s", movieID.String(), filepath.Ext(file.Filename))
	filePath := filepath.Join(dir, fileName)

	// Save the file
	if err := saveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// saveUploadedFile saves the uploaded file to the specified path
func saveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
