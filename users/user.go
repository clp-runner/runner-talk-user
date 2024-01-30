package users

import "database/sql"

type User struct {
	Provider          string 			`json:"provider" db:"provider"`
	UserID            string 			`json:"user_id" db:"user_id"`
	UserName          string 			`json:"user_name" db:"user_name"`
	Email             string 			`json:"email" db:"email"`
	DisplayName       sql.NullString 	`json:"display_name" db:"display_name"`
	ProfilePictureUrl sql.NullString 	`json:"profile_picture_url" db:"profile_picture_url"`
	Token             sql.NullString 	`json:"token" db:"token"`
	ExpirationTime    sql.NullString 	`json:"expiration_time" db:"expiration_time"`
	CreatedAt         sql.NullString 	`json:"created_at" db:"created_at"`
	LastLoginAt       sql.NullString 	`json:"last_login_at" db:"last_login_at"`
}

type UserInterface interface {
	Login()
	Callback()
	Logout()
}