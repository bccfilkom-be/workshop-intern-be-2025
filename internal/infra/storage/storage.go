package storage

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"mime/multipart"
)

type storage struct {
	client *supabasestorageuploader.Client
}

type Service interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	DeleteFile(link string) error
}

func New() Service {
	supClient := supabasestorageuploader.New(
		"https://ffegebxperqkyegcldbf.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZmZWdlYnhwZXJxa3llZ2NsZGJmIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc0MDg5MzAxMCwiZXhwIjoyMDU2NDY5MDEwfQ.pQE2GnzLSs75QvycwG-Z9RBdgRqfFoccSaQ-Ghdk0Eg",
		"akbar-bucket",
	)
	return storage{
		client: supClient,
	}
}

func (s storage) UploadFile(file *multipart.FileHeader) (string, error) {
	return s.client.Upload(file)
}

func (s storage) DeleteFile(link string) error {
	return s.client.Delete(link)
}
