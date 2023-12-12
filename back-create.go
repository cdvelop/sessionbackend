package sessionbackend

import (
	"fmt"

	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
	"github.com/cdvelop/token"
)

func (s sessionBackend) Create(u *model.User, params ...map[string]string) (err string) {
	const this = "Create session error "

	data_db, err := s.Checking(u, params)
	if err != "" {
		return this + err
	}

	// 1- CREAMOS EL OBJETO USUARIO CON SU TOKEN

	key_area := data_db[s.FieldArea]

	new_user := model.User{
		Token:          token.BuildUniqueKey(16),
		Id:             data_db[s.FieldID],
		Ip:             u.Ip,
		Name:           data_db[s.FieldName],
		Area:           key_area,
		AreaName:       s.Config.AreasName[key_area],
		AccessLevel:    data_db[s.FieldAccessLevel],
		LastConnection: s.DateToDayHour(),
	}

	// fmt.Println("\nUSUARIO:", new_user)

	// 2- CONVERTIMOS LA DATA EN BYTES JSON
	encode_user, err := s.EncodeStruct(new_user)
	if err != "" {
		return this + err
	}

	fmt.Println("3- CIFRAMOS LA DATA DEL USUARIO:")
	session_encode, err := s.CipherAdapter.Encrypt(encode_user)
	if err != "" {
		return this + err
	}

	fmt.Println("4- CREAMOS EL OBJETO SESIÓN DEL LADO DEL CLIENTE")

	new_session := sessionhandler.SessionStore{
		Id_session:     new_user.Id,
		Session_status: "in",
		Session_encode: session_encode,
	}

	//5- CONVERTIMOS A JSON LA SESIÓN
	encode_session, err := s.EncodeStruct(new_session)
	if err != "" {
		return this + err
	}

	//6- CREAMOS UN NUEVO MAPA CON LA NUEVA SALIDA DE INFORMACIÓN
	response := map[string]string{
		"session": string(encode_session),
	}

	// fmt.Println("\nnew_user:", new_user)

	out, err := s.BackendLoadBootData(&new_user)
	if err != "" {

		return this + err
	}
	response["boot"] = out

	//7- REMPLAZAMOS EL PRIMER ELEMENTO CON LA NUEVA INFORMACIÓN
	params[0] = response

	// fmt.Println("DATA ENVIADA:", params)

	return
}
