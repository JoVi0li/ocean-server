package util

import "errors"

var ErrorUnknownClaimsType = errors.New("unknown token claims type, cannot proceed")