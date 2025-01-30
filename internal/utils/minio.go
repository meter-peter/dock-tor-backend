package utils

import (
	"context"
	"log"

	"clinic-management/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func InitMinIO() {
	cfg := config.GetConfig()

	// Initialize minio client
	c, err := minio.New(cfg.MINIOEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			cfg.MINIOAccessKey,
			cfg.MINIOSecretKey,
			"",
		),
		Secure: false,
	})

	if err != nil {
		log.Fatalf("MinIO connection failed: %v", err)
	}

	// Create bucket if not exists
	exists, err := c.BucketExists(context.Background(), "patients")
	if err == nil && !exists {
		err = c.MakeBucket(context.Background(), "patients", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	}

	client = c
}

func GetMinIOClient() *minio.Client {
	return client
}
