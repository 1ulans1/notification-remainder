package telegram

type TgService interface {
	Send(destination string, text string) error
}

type TgServiceImpl struct {
}
