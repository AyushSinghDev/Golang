//When to use pointers?
//When we need to update
//When we need to call the object a lot of time of some reasonable size

package main

import "fmt"

type User struct {
	email    string
	password string
	age      int
}

func (u User) Email() string {
	return u.email
}

//A pointer enables indirect access and modification of the values
func (u *User) udpateEmail(email string) {
	u.email = email
}

func main() {
	user := User{
		email: "ayush@gmail.com",
	}
	user.udpateEmail("devxt@hotmail.com")
	fmt.Println(user.Email())

}
