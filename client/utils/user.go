package utils

import "os/user"

//GetUsername retrives user name
func GetUsername() string {
	usr, _ := user.Current()
	return usr.Username
}
