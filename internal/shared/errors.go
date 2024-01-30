package shared

import "errors"

var ErrorCredentialsInvalid = errors.New("email or password invalid")
var ErrorUnknownClaimsType = errors.New("unknown token claims type, cannot proceed")
var ErrorExpiredToken = errors.New("token expired")
var ErrorMissingAuthorizationToken = errors.New("missing the authorization token")
var ErrorInvalidAuthorizationToken = errors.New("missing the authorization token")
var ErrorUsernameInvalid = errors.New("username empty or invalid")
var ErrorEmailInvalid = errors.New("email empty or invalid")
var ErrorPasswordInvalid = errors.New("password empty or invalid")
var ErrorUserNotFound = errors.New("user not found")
var ErrorTryingHashPassword = errors.New("error trying hash the password")
var ErrorIdInvalid = errors.New("id empty or invalid")