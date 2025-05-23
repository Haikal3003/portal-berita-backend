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

func (s *CloudinaryService) UploadThumbnailImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	uploadResult, err := s.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "thumbnail_berita",
	})

	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
