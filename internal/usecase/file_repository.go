package usecase

import "miniDriver/internal/model"

type FileRepository interface {
	SaveFile(file model.File) error
	ListFiles() ([]model.File, error)
	DeleteFile(filename string) error
	FileExists(filename string) bool
}
