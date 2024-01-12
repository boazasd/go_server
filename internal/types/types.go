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

type AgoraAgent struct {
	Id          int64
	UserId      int64     `db:"userId"`
	UserEmail   string    `db:"userEmail"`
	SearchTxt   string    `db:"searchTxt"`
	Category    string    `db:"category"`
	SubCategory string    `db:"subCategory"`
	Area        string    `db:"area"`
	Condition   string    `db:"condition"`
	WithImage   bool      `db:"withImage"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

type AgoraData struct {
	Id             int64     `db:"id"`
	Link           string    `db:"link"`
	Name           string    `db:"name"`
	Details        string    `db:"details"`
	City           string    `db:"city"`
	Category       string    `db:"category"`
	MiddleCategory string    `db:"middleCategory"`
	SubCategory    string    `db:"subCategory"`
	Condition      string    `db:"condition"`
	Area           string    `db:"area"`
	Image          string    `db:"image"`
	Processed      bool      `db:"processed"`
	Date           time.Time `db:"date"`
	CreatedAt      time.Time `db:"createdAt"`
	UpdatedAt      time.Time `db:"updatedAt"`
}

type AgoraAgentResults struct {
	UserId  int64  `db:"userId"`
	AgentId int64  `db:"agentId"`
	Email   string `db:"userEmail"`
	AgoraData
}

type Session struct {
	Id             int64
	SessionId      string
	UserId         int64
	ExpirationTime time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
