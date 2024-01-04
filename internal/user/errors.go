package user

import "errors"

var ErrorUsernameInvalid = errors.New("username empty or invalid")
var ErrorEmailInvalid = errors.New("email empty or invalid")
var ErrorPasswordInvalid = errors.New("password empty or invalid")
var ErrorUserNotFound = errors.New("user not found")
var ErrorTryingHashPassword = errors.New("error trying hash the password")