package service

import (
	"fmt"
	"time"

	"github.com/raaaaaaaay86/time-now-go-unit-test/model"
)

type IMessageService interface {
	Send(title, content string) model.MessageInfo
}

type MessageService struct {
}

func (svc MessageService) Send(title string, content string) (model.MessageInfo, error) {
	msg := model.MessageInfo{
		Title:    title,
		Content:  content,
		SendTime: time.Now(),
	}

	fmt.Println("Send message...")

	return msg, nil
}
