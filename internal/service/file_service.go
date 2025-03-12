package service

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"miniDriver/internal/model"
	"miniDriver/internal/usecase"
	"os"
	"path/filepath"
	"time"
)

type FileService struct {
	repo      usecase.FileRepository
	uploadDir string
	shared    map[string]string
}

func NewFileService(repo usecase.FileRepository, uploadDir string) *FileService {
	os.Mkdir(uploadDir, os.ModePerm)
	return &FileService{repo: repo, uploadDir: uploadDir, shared: make(map[string]string)}
}

func (s *FileService) UploadFile(filename string, size int64, file io.Reader) error {
	if s.repo.FileExists(filename) {
		return errors.New("файл уже существует")
	}

	dstPath := filepath.Join(s.uploadDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return s.repo.SaveFile(model.File{Name: filename, Size: size})
}

func (s *FileService) ListFiles() ([]model.File, error) {
	return s.repo.ListFiles()
}

func (s *FileService) DeleteFile(filename string) error {
	if err := s.repo.DeleteFile(filename); err != nil {
		return err
	}

	filePath := filepath.Join(s.uploadDir, filename)
	return os.Remove(filePath)
}

func (s *FileService) ShareFile(filename string) (string, error) {
	if !s.repo.FileExists(filename) {
		return "", errors.New("файл не найден")
	}

	linkID := generateUUID()
	s.shared[linkID] = filename

	return "http://localhost:8080/public/" + linkID, nil
}

func generateUUID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}
