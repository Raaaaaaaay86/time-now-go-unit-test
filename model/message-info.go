package model

import "time"

type MessageInfo struct {
	Title    string
	Content  string
	SendTime time.Time
}
