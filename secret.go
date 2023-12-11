package sessionbackend

import "github.com/cdvelop/token"

var secret_key_signing = ""

func AddPrivateSecretKeySigning() map[string]string {
	return map[string]string{
		`github.com/cdvelop/sessionbackend.secret_key_signing`: token.BuildUniqueKey(32),
	}
}
