package sessionbackend

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
	"github.com/cdvelop/timetools"
	"github.com/cdvelop/token"
)

func (s sessionBackend) Create(u *model.User, params ...map[string]string) (err string) {
	const e = "Create session "

	data_db, err := s.Checking(u, params)
	if err != "" {
		return e + err
	}

	var w http.ResponseWriter

	if rw, ok := u.W.(http.ResponseWriter); ok {
		w = rw
	}

	if w == nil {
		return e + "parámetro http.ResponseWriter incorrecto"
	}

	// 0-

	// 1- CREAMOS EL OBJETO USUARIO CON SU TOKEN

	key_area := data_db[s.FieldArea]

	date, hour := s.DateToDayHour(&timetools.DateFormat{
		LeftDay:     false,
		WithSeconds: true,
	})

	new_user := model.User{
		Token:          token.BuildUniqueKey(16),
		Id:             data_db[s.FieldID],
		Ip:             u.Ip,
		Name:           data_db[s.FieldName],
		Area:           key_area,
		AreaName:       s.Config.AreasName[key_area],
		AccessLevel:    data_db[s.FieldAccessLevel],
		LastConnection: date + " " + hour,
	}

	fmt.Println("\nUSUARIO:", new_user)

	// 2- CONVERTIMOS LA DATA EN BYTES JSON
	encode_user, err := s.EncodeStruct(new_user)
	if err != "" {
		return e + err
	}

	// fmt.Println("3- CIFRAMOS LA DATA DEL USUARIO.")
	session_encode, err := s.CipherAdapter.Encrypt(encode_user)
	if err != "" {
		return e + err
	}

	// fmt.Println("4- CREAMOS LA COOKIE DE SESSION")
	s.Gookie.Set(session_encode, w)

	session_number, its_new_number, err := s.getSessionNumber(new_user.Id)
	if err != "" {
		return e + err
	}

	// fmt.Println("5- CREAMOS EL OBJETO SESIÓN DEL LADO DEL CLIENTE")
	response := map[string]string{
		s.Id_session:     new_user.Id,
		s.Session_number: session_number,
		s.Session_encode: session_encode,
	}

	// ALMACENAMOS LA SESIÓN EN EL BACKEND SOLO SI ES UN NUMERO NUEVO
	if its_new_number {
		err = s.CreateObjectsInDB(s.Table, false, response)
		if err != "" {
			return e + err
		}
	}

	// fmt.Println("\nnew_user:", new_user)

	//7- AGREGAMOS DATA DE ARRANQUE
	boot_data := s.BackendLoadBootData(&new_user).JsonBootActions
	// fmt.Println("boot_data:", boot_data)
	if boot_data != "none" {
		response["boot"] = boot_data
	}

	//8- REMPLAZAMOS EL PRIMER ELEMENTO CON LA NUEVA INFORMACIÓN
	params[0] = response

	return
}
