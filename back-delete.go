package sessionbackend

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (s sessionBackend) Delete(u *model.User, data ...map[string]string) (err string) {

	fmt.Println("ELIMINANDO SESSION USUARIO:", u)

	return
}
