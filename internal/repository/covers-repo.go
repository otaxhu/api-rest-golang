package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type CoversRepository interface {
	SaveMovieCover(cover *multipart.FileHeader) error
	DeleteCover(coverUrl string) error
}

type coversRepoImpl struct{}

func NewCoversRepository() CoversRepository {
	return &coversRepoImpl{}
}

func (repo *coversRepoImpl) SaveMovieCover(cover *multipart.FileHeader) error {
	file, err := cover.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	path, err := filepath.Abs("./static/covers")
	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(path, cover.Filename))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return err
}

func (repo *coversRepoImpl) DeleteCover(coverUrl string) error {
	filename := filepath.Base(coverUrl)
	path, err := filepath.Abs("./static/covers")
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(path, filename))
}
