package sessionbackend

import (
	"strconv"

	"github.com/cdvelop/model"
)

// en el backend no se requiere agregar numero de usuario al crear un nuevo id
func (s sessionBackend) UserSessionNumber() (number string, err string) {
	return "", ""
}

// findNextNumber encuentra el siguiente número correlativo disponible
func (s sessionBackend) getSessionNumber(session_id string) (number string, its_new_number bool, err string) {

	const e = "getSessionNumber "

	result, err := s.ReadSyncDataDB(&model.ReadParams{FROM_TABLE: s.Table, ORDER_BY: s.Session_number})
	if err != "" {
		err = e + err
		return
	}

	// Crear un mapa para rastrear los números utilizados
	usedNumbers := make(map[int]bool)

	// Iterar sobre el resultado y marcar los números utilizados
	for _, row := range result {

		if id, id_exist := row[s.Id_session]; id_exist && session_id == id {
			return row[s.Session_number], false, ""
		}

		if strNumber, ok := row[s.Session_number]; ok {

			intNumber, er := strconv.Atoi(strNumber)
			if er != nil {
				err = e + er.Error()
				return
			}
			usedNumbers[intNumber] = true
		}
	}

	// Encontrar el siguiente número correlativo disponible
	for i := 1; ; i++ {
		if !usedNumbers[i] {
			return strconv.Itoa(i), true, ""
		}
	}

}
