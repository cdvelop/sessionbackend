package sessionbackend

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.Handlers) (b *sessionBackend, err string) {

	s, err := sessionhandler.Add(h)
	if err != "" {
		return nil, err
	}

	b = &sessionBackend{
		Session: s,
	}

	h.AuthAdapter = b

	return b, ""
}
