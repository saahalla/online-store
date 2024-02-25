package dto

type UserDB struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	DefaultDate
}

type StructNull struct {
}
