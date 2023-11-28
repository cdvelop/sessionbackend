package sessionbackend

import (
	"net/http"

	"github.com/cdvelop/model"
)

func (s sessionBackend) GetLoginUser(http_request any) (user *model.User, err string) {
	var r *http.Request

	if rq, ok := http_request.(*http.Request); ok {
		r = rq
	}

	if r == nil {
		return nil, "error GetLoginUser. parámetro *http.Request no en enviado en authbackend"
	}

	// if !s.production_mode {

	// 	user = &model.User{
	// 		Token:          "123",
	// 		Id:             "1635572582072481400",
	// 		Ip:             "172.0.0.1", //"172.0.0.41"
	// 		Name:           "dra. karla acero",
	// 		Area:           "s",
	// 		AccessLevel:    "",
	// 		LastConnection: "",
	// 	}
	// 	return user, ""
	// }

	return nil, "usuario no registrado"
}
