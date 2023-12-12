package sessionbackend

import (
	"net/http"

	"github.com/cdvelop/model"
)

func (s sessionBackend) BackendCheckUser(http_request any) (user *model.User, err string) {
	const this = "BackendCheckUser error "
	var r *http.Request

	if rq, ok := http_request.(*http.Request); ok {
		r = rq
	}

	if r == nil {
		return nil, this + "parámetro *http.Request incorrecto"
	}

	// 1
	cookie, err := s.Gookie.Get(r)
	if err != "" {
		return nil, "usuario no registrado"
	}

	// fmt.Println("header:", r.Header)

	// fmt.Println("cookie value", cookie.Value)

	return s.DecodeUser(cookie.Value)
}
