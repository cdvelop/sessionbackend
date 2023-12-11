package sessionbackend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cdvelop/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
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
	oauth2_token, err := s.getTokenFromCookie(r)
	if err != "" {
		return nil, this + err
	}

	fmt.Println("oauth2_token:", oauth2_token)

	// if !s.production_mode {

	// 	user = &model.User{
	// 		Token:          "123",
	// 		Id:             "1635572582072481400",
	// 		Ip:             "172.0.0.1", //"172.0.0.41"
	// 		Name:           "drs. karla acero",
	// 		Area:           "s",
	// 		AccessLevel:    "",
	// 		LastConnection: "",
	// 	}
	// 	return user, ""
	// }

	return nil, "usuario no registrado"
}

func (s sessionBackend) getTokenFromCookie(r *http.Request) (out *oauth2.Token, err string) {

	const this = "error token "

	// 1- OBTENGO LA COOKIE
	cookie, err := s.Gookie.Get(r)
	if err != "" {
		return nil, this + err
	}

	// 2- Obtén el valor del token JWT de la cookie
	tokenString := cookie.Value

	// 3- Analiza y verifica el token JWT
	token, er := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Devuelve la misma clave secreta utilizada para firmar el token
		return []byte(secret_key_signing), nil
	})
	if er != nil {
		err = this + "no se logro verificar"
		s.Log(err + " " + er.Error())
		return
	}

	// Verifica si el token es válido
	if token.Valid {
		// Accede a las reclamaciones del token según tus necesidades
		claims := token.Claims.(jwt.MapClaims)
		// Construye un *oauth2.Token a partir de las reclamaciones del token
		accessToken := claims["AccessToken"].(string)
		tokenType := claims["TokenType"].(string)
		refreshToken := claims["RefreshToken"].(string)
		expiry := claims["Expiry"].(string) // Asumiendo que la reclamación Expiry está en formato string

		expiryTime, er := time.Parse(time.RFC3339, expiry)
		if er != nil {
			err = this + "expirado"
			s.Log(err + " " + er.Error())
			return
		}

		token := &oauth2.Token{
			AccessToken:  accessToken,
			TokenType:    tokenType,
			RefreshToken: refreshToken,
			Expiry:       expiryTime,
		}

		return token, ""
	}

	return nil, this + "no válido"

}
