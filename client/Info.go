package main

import "os/user"

func GetUsername() string {
	usr, _ := user.Current()
	return usr.Username
}

/*func GetAV() string {
	res := Run("WMIC /Node:localhost /Namespace:\\root\\SecurityCenter2 Path AntiVirusProduct Get displayName /Format:List")
	return res
}
*/
