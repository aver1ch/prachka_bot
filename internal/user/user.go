package user

type User struct {
	userName string
	room     int
	chatID   int64
}

var Users []User
