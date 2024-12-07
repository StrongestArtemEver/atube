package services

type PingService struct{}

// @Summary      Ping endpoint
// @Description  Проверяет работоспособность сервера
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func(ps *PingService) GetPingMessage() string{
	return "pong"
}