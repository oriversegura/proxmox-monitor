package models

import "time"

type status string

const (
	StatusOnline  status = "Online"
	StatusOffline status = "Offline"
)


type Base struct {
	Name      string
	IP        string
	Timestamp time.Time
	Status status
}