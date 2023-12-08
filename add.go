package sessionbackend

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.Handlers, c sessionhandler.Config) (b *sessionBackend, err string) {

	s, err := sessionhandler.Add(h, c)
	if err != "" {
		return nil, err
	}

	b = &sessionBackend{
		Session: s,
	}

	h.AuthBackendAdapter = b

	return b, ""
}
