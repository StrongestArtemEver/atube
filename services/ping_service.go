package services

type PingService struct{}

func(ps *PingService) GetPingMessage() string{
	return "pong"
}