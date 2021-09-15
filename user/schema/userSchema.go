package userSchema

import "fmt"



type User struct {
	Id       int64
	Name     string
	Password string
	Email    string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
}
