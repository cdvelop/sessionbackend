package sessionbackend

import "github.com/cdvelop/token"

const secret_key_signing = ""

func AddPrivateSecretKeySigning(new_key string) map[string]string {
	return map[string]string{
		`github.com/cdvelop/sessionbackend.secret_key_signing`: token.BuildUniqueKey(32),
	}
}
