package util

import "errors"

var ErrorUnknownClaimsType = errors.New("unknown token claims type, cannot proceed")
var ErrorExpiredToken = errors.New("token expired")
var ErrorMissingAuthorizationToken = errors.New("missing the authorization token")
