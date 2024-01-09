package users

type User struct {
	userID 				string
	userName 			string
	email 				string
	displayName 		string
	profilePictureUrl 	string
	token 				string
	expirationTime 		string
	createdAt 			string
	lastLoginAt 		string
}

type UserAction interface {
	Login()
	Callback()
	Logout()
}