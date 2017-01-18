package main

import "os/user"

func GetUsername() string {
	usr, _ := user.Current()
	return usr.Username
}