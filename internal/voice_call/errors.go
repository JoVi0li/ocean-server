package voice_call

import "errors"

var ErrorMissingParticipant = errors.New("missing participant")
var ErrorParticipantInvalid = errors.New("participant invalid")
var ErrorVoiceCallNotFound = errors.New("voice call not found")