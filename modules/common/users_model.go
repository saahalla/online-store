package common

type UserDB struct {
	DefaultDate
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
}

type StructNull struct {
}
