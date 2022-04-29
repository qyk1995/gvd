package utils

var (
	RegistuserVerify = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}}
	LoginVerify      = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}}
)
