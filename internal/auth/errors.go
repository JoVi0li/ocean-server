package auth

import "errors"

var ErrorCredentialsInvalid = errors.New("email or password invalid")
var ErrorMissingAuthorizationToken = errors.New("missing the authorization token")
