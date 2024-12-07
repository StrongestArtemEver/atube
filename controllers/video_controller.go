package controllers

import (
	"atube/services"
	"context"
	"fmt"
	"net/http"

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

// GetAllVideo получает список всех видео.
// @Summary Получить список видео
// @Description Возвращает список всех доступных видео
// @Tags Video
// @Produce json
// @Success 200 {array} models.Video "Успешный ответ"
// @Router /videos [get]
func (vc *VideoController) GetAllVideo(c *gin.Context) {
	ctx := context.Background()

	videos, err := vc.videoService.GetAllVideo(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videos})
}
