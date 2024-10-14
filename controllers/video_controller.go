package controllers

import (
	"atube/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VideoController struct {
	videoService *services.VideoService
}

func NewVideoController(db *gorm.DB) (*VideoController, error) {
	videoService, err := services.NewVideoService(db)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать VideoService: %v", err)
	}

	return &VideoController{
		videoService: videoService,
	}, nil
}

func (vc *VideoController) UploadVideo(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "Только метод POST поддерживается"})
		return
	}

	videoURL, thumbnailURL, err := vc.videoService.UploadVideo(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := map[string]string{
		"message":       "Видео успешно загружено",
		"video_url":     videoURL,
		"thumbnail_url": thumbnailURL,
	}

	c.JSON(200, response)
}
