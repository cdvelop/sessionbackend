package sessionbackend

import "github.com/cdvelop/model"

func Add(h *model.Handlers) (s *sessionBackend, err string) {

	s = &sessionBackend{}

	h.AuthAdapter = s

	return s, ""
}
