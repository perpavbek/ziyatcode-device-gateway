package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileService struct {
	RootDir string
}

func NewFileService(rootDir string) *FileService {
	return &FileService{RootDir: rootDir}
}

func (s *FileService) SaveFile(filePath string, fileName string, data []byte) (string, error) {
	dirPath := s.resolveStatic(filePath)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("saving file error: %w", err)
	}

	savePath := filepath.Join(dirPath, fileName)
	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return "", fmt.Errorf("saving file error: %w", err)
	}

	return filepath.Join(filePath, fileName), nil
}

func (s *FileService) DeleteFile(path string) error {
	filePath := s.resolveStatic(path)
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("deleting file error: %w", err)
	}
	return nil
}

func (s *FileService) GetFile(path string) ([]byte, error) {
	filePath := s.resolveStatic(path)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

func (s *FileService) resolveStatic(dirPath string) string {
	return filepath.Join(s.RootDir, dirPath)
}