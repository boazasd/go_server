package types

import "time"

type User struct {
	Id        int64
	FirstName string `db:"firstName"`
	LastName  string `db:"lastName"`
	Email     string
	Password  string
	Roles     string
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

type Session struct {
	Id             string
	SessionId      string
	UserId         int64
	ExpirationTime time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type AgoraData struct {
	Link      string
	Name      string
	Details   string
	City      string
	Area      string
	State     string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
