package sessionbackend

import (
	"net/http"

	"github.com/cdvelop/model"
	"github.com/dgrijalva/jwt-go"
)

func (s *sessionBackend) setTokenInCookie(user *model.User, r *http.Request, w http.ResponseWriter) (err string) {

	//1- Genera un token JWT con la información del usuario
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		s.Gookie.Name: user,
	})

	//2- Firma el token con una clave secreta
	tokenString, er := jwtToken.SignedString([]byte(secret_key_signing))
	if er != nil {
		return "error al generar el token: " + er.Error()
	}

	//3- setear la cookie
	s.Gookie.Set(tokenString, w)

	return ""
}
