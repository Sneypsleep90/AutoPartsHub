package models

import "time"

type User struct {
	userId    string
	firstName string
	lastName  string
	balance   int
	createdAt time.Time
	isActive  bool
}
