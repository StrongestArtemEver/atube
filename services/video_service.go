package services

import (
	"atube/models"
	"atube/storage"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoService struct {
	db        *gorm.DB
	s3Storage *storage.S3Storage
}

func NewVideoService(db *gorm.DB) (*VideoService, error) {
	ctx := context.Background()
	bucket := "firstartemsb"

	s3Storage, err := storage.NewS3Storage(ctx, bucket)
	if err != nil {
		return nil, fmt.Errorf("не удалось инициализировать S3Storage: %v", err)
	}

	return &VideoService{
		db:        db,
		s3Storage: s3Storage,
	}, nil
}

func (vs *VideoService) UploadVideo(c *gin.Context) (string, string, error) {
	videoFile, videoHeader, err := c.Request.FormFile("video")
	if err != nil {
		return "", "", fmt.Errorf("не удалось получить видеофайл: %v", err)
	}
	defer videoFile.Close()

	thumbnailFile, thumbnailHeader, err := c.Request.FormFile("thumbnail")
	if err != nil {
		return "", "", fmt.Errorf("не удалось получить миниатюру: %v", err)
	}
	defer thumbnailFile.Close()

	videoKey := uuid.New().String() + "_" + videoHeader.Filename
	thumbnailKey := uuid.New().String() + "_" + thumbnailHeader.Filename

	ctx := context.Background()

	videoURL, err := vs.s3Storage.UploadFile(ctx, videoKey, videoFile)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при загрузке видеофайла: %v", err)
	}

	thumbnailURL, err := vs.s3Storage.UploadFile(ctx, thumbnailKey, thumbnailFile)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при загрузке миниатюры: %v", err)
	}

	video := &models.Video{
		UserID:       1, // Предполагается, что пользователь авторизован
		Title:        c.PostForm("title"),
		Description:  c.PostForm("description"),
		FileURL:      videoURL,
		ThumbnailURL: thumbnailURL,
		Duration:     300, // Длительность должна быть определена автоматически
	}

	if err := models.SaveVideo(vs.db, video); err != nil {
		return "", "", fmt.Errorf("ошибка при сохранении видео в базе данных: %v", err)
	}

	return videoURL, thumbnailURL, nil
}


func (vs *VideoService) GetAllVideo(ctx context.Context) ([]models.Video,error) {
	var videos []models.Video

	if err := vs.db.Find(&videos).Error; err != nil {
		return nil,fmt.Errorf("ошибка при получении видео: %v" ,err)
	}
	
	return videos,nil
}