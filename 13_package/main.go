package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/geekAshish/go-programming/tree/main/13_package/auth"
	"github.com/geekAshish/go-programming/tree/main/13_package/user"
)

func main() {
	auth.LoginWithCredentials("ashish", "kushwaha")
	sessionDetail := auth.GetSession()
	
	user := user.User{
		Email: "ashish@cool.com",
		Username: "kakoii",
	}

	fmt.Println(sessionDetail, user)

	color.Cyan(user.Username)
	color.Red(user.Email)

}
