package service

import (
	"fmt"
	"testing"
	"time"
)

func TestMessageService_Send(t *testing.T) {
	svc := MessageService{}
	expectedSendTime := time.Now()
	expectedTitle := "The output MessageInfo.Title should be this."
	expectedContent := "The output MessageInfo.Content should be this."

	msg, err := svc.Send(expectedTitle, expectedContent)
	if err != nil {
		t.Fatalf("got error when testing MessageService.Send(): %s", err)
	}

	assertEqual(t, "expected title and actual title not same", expectedTitle, msg.Title)
	assertEqual(t, "expected content and actual content not same", expectedContent, msg.Content)
	assertEqual(t, "expected send time and actual send time not same", expectedSendTime, msg.SendTime)
}

func TestBetterMessageService_Send(t *testing.T) {
	now := time.Now()
	svc := BetterMessageService{
		Now: func() time.Time {
			return now
		},
	}
	expectedSendTime := now
	expectedTitle := "The output MessageInfo.Title should be this."
	expectedContent := "The output MessageInfo.Content should be this."

	msg, err := svc.Send(expectedTitle, expectedContent)
	if err != nil {
		t.Fatalf("got error when testing MessageService.Send(): %s", err)
	}

	assertEqual(t, "expected title and actual title not same", expectedTitle, msg.Title)
	assertEqual(t, "expected content and actual content not same", expectedContent, msg.Content)
	assertEqual(t, "expected send time and actual send time not same", expectedSendTime, msg.SendTime)
}

func assertEqual(t *testing.T, message string, expected, actual any) {
	if expected != actual {
		equalFail(t, message, expected, actual)
	}
}

func equalFail(t *testing.T, message string, expected, actual any) {
	fmt.Println("===")
	{
		fmt.Printf("\033[37;102mExpected: %v\033[0m\n", expected)
		fmt.Printf("\033[37;101mActual: %v\033[0m\n", actual)
	}
	fmt.Println("===")

	t.Fatal(message)
}
