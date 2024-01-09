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
	Id             int64
	Link           string
	Name           string
	Details        string
	City           string
	Category       string `db:"category"`
	MiddleCategory string `db:"middleCategory"`
	SubCategory    string `db:"subCategory"`
	Condition      string `db:"condition"`
	Area           string
	Image          string
	Processed      bool
	Date           time.Time
	CreatedAt      time.Time `db:"createdAt"`
	UpdatedAt      time.Time `db:"updatedAt"`
}

type AgoraAgentResults struct {
	UserId  int64  `db:"userId"`
	AgentId int64  `db:"agentsId"`
	Email   string `db:"email"`
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
