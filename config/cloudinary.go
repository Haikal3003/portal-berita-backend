package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUD_NAME"),
		os.Getenv("CLOUD_API_KEY"),
		os.Getenv("CLOUD_SECRET_KEY"),
	)

	if err != nil {
		log.Fatalf("Cloudinary config error: %v", err)
	}

	return cld
}
