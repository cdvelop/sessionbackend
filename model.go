package sessionbackend

import (
	"github.com/cdvelop/sessionhandler"
)

type sessionBackend struct {
	*sessionhandler.Session
}
