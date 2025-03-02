package usecase

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/storage"
	"mime/multipart"
)

type FileUsecaseItf interface {
	Upload(file *multipart.FileHeader) (string, error)
}

type FileUsecase struct {
	storage storage.Service
}

func NewFileUsecase(storage storage.Service) FileUsecaseItf {
	return &FileUsecase{
		storage: storage,
	}
}

func (u *FileUsecase) Upload(file *multipart.FileHeader) (string, error) {
	return u.storage.UploadFile(file)
}
