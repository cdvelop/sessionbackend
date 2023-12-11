package sessionbackend

import (
	"github.com/cdvelop/gookie"
	"github.com/cdvelop/sessionhandler"
)

type sessionBackend struct {
	*sessionhandler.Session
	*gookie.Gookie
}
