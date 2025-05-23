package services

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService(cld *cloudinary.Cloudinary) *CloudinaryService {
	return &CloudinaryService{
		cld: cld,
	}
}

func (s *CloudinaryService) UploadImage(fileHeader *multipart.FileHeader, folder string) (string, string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}

	defer file.Close()

	uploadResult, err := s.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: folder,
	})

	if err != nil {
		return "", "", err
	}

	return uploadResult.SecureURL, uploadResult.PublicID, nil
}

func (s *CloudinaryService) DeleteImage(publicID string) error {
	_, err := s.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})

	return err
}
