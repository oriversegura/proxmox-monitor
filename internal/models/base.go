package models

import "time"

type Status string

const (
	StatusOnline  Status = "Online"
	StatusOffline Status = "Offline"
)

type Base struct {
	Name      string
	IP        string
	Timestamp time.Time
	Status    Status
}
