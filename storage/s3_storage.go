package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Storage struct{
	Client *s3.Client
	Bucket string
}

func NewS3Storage(ctx context.Context, bucket string ) (*S3Storage,error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Файл .env не найден")
}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil,fmt.Errorf("не удалось загрузить конфигурацию AWS: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Storage{
		Client: client,
		Bucket: bucket,
	},nil
}


func (s *S3Storage) UploadFile(ctx context.Context, key string, file io.Reader) (string, error) {
	_, err := s.Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(key),
			Body:   file,
	})
	if err != nil {
			return "", fmt.Errorf("ошибка при загрузке файла: %v", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.Bucket, key)
	return url, nil
}

func (s *S3Storage) DownloadFile(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при скачивании файла: %v", err)
	}
	return output.Body, nil
}

func (s *S3Storage) ListObjects(ctx context.Context, prefix string) ([]types.Object, error) {
	output, err := s.Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка объектов: %v", err)
	}
	return output.Contents, nil
}

func (s *S3Storage) DeleteObject(ctx context.Context, key string) error {
	_, err := s.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("ошибка при удалении объекта: %v", err)
	}
	return nil
}