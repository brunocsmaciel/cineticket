package types

import (
	"time"
)

type Movie struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	RunningTime  int       `json:"running_time"`
	Actors       []string  `json:"actors"`
	Languages    []string  `json:"languages"`
	Genre        []string  `json:"genre"`
	PremiereDate time.Time `json:"premiere_date"`
}
