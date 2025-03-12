package usecase

import (
	"io"
	"miniDriver/internal/model"
)

type FileUseCase interface {
	UploadFile(filename string, size int64, file io.Reader) error
	ListFiles([]model.File, error)
	DeleteFiles(filename string) error
	ShareFile(filename string) (string, error)
}
