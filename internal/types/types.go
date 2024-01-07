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

type Wishes struct {
	Id        int64
	UserId    int64 `db:"userId"`
	Wishes    string
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

type Session struct {
	Id             int64
	SessionId      string
	UserId         int64
	ExpirationTime time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type AgoraData struct {
	Id        int64
	Link      string
	Name      string
	Details   string
	City      string
	Area      string
	State     string
	Date      time.Time
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}
