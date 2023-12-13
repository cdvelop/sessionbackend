package sessionbackend

import (
	"github.com/cdvelop/model"
)

func (c sessionBackend) BackendLoadBootData(u *model.User) (out model.BootPageData) {
	const t = "BackendLoadBootData error :"

	out = model.BootPageData{
		JsonBootActions: "none",
	}

	if u == nil {
		c.Log(t + "no se puede entregar información si no te has registrado")
		return
	}

	var responses []model.Response
	for _, o := range c.GetObjects() {

		// fmt.Println("BackHandler.BootResponse", o.ObjectName)
		// fmt.Println("Estado Back:", o.BackHandler.BootResponse)

		if o.BackHandler.BootResponse != nil {
			resp, err := o.BackHandler.AddBootResponse(u)
			if err != "" {
				c.Log(t + o.ObjectName + " " + err)
				return
			} else if len(resp) != 0 {
				responses = append(responses, resp...)
			}
		}
	}

	boot_data_byte, err := c.EncodeResponses(responses...)
	if err != "" {
		c.Log(t + err)
		return
	}

	out.JsonBootActions = string(boot_data_byte)

	return
}
