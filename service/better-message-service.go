package service

import (
	"fmt"
	"time"

	"github.com/raaaaaaaay86/time-now-go-unit-test/model"
)

type IBetterMessageService interface {
	Send(title, content string) model.MessageInfo
}

type BetterMessageService struct {
	Now func() time.Time
}

func (svc BetterMessageService) Send(title string, content string) (model.MessageInfo, error) {
	msg := model.MessageInfo{
		Title:    title,
		Content:  content,
		SendTime: svc.Now(),
	}

	fmt.Println("Send message...")

	return msg, nil
}
