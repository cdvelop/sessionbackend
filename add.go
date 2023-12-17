package sessionbackend

import (
	"github.com/cdvelop/gookie"
	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.MainHandler, sc *sessionhandler.Config) (err string) {

	sh, err := sessionhandler.Add(h, sc)
	if err != "" {
		return err
	}

	sb := &sessionBackend{
		Session: sh,
		Gookie: &gookie.Gookie{
			Name:             sessionhandler.TABLE_NAME,
			Domain:           "",
			Https:            false,
			CookieExpiration: sc.CookieExpiration,
		},
	}

	h.SessionBackendAdapter = sb
	h.BackendBootDataUser = sb

	sh.Form.BackHandler.CreateApi = sb
	sh.Form.BackHandler.DeleteApi = sb

	h.CreateTablesInDB([]*model.Object{sh.Object}, func(err string) {
		if err != "" {
			h.Log("AddAuthAdapter sessionbackend " + err)
		}
	})

	return
}
