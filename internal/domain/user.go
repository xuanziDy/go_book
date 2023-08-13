package domain

import "time"

type User struct {
	Id           int64
	Email        string
	Password     string
	Nickname     string `json:"nickname"`
	Birthday     string `json:"birthday"`
	Introduction string `json:"introduction"`
	Ctime        time.Time
}
