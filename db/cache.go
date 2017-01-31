package db

var accounts map[string]*User

func CreateAccount(user *User) {
	accounts[user.Email] = user
}
