package entities

import "time"

type Login struct {
	UserID int64
	Hash   string
}

type Session struct {
	Token       string
	UserId      int64
	CreateTime  time.Time
	TTL         int
	Fingerprint Fingerprint
}

type Fingerprint struct {
	Ip      string
	Browser string
}
