package authcode

import (
	"github.com/zalando/go-keyring"
)

const (
	service = "stravastats"
	user    = "authcode"
)

func ReadFromStore() (string, error) {
	code, err := keyring.Get(service, user)
	if err != nil {
		return "", err
	}

	return code, nil
}

func SaveToStore(code string) error {
	err := keyring.Set(service, user, code)
	return err
}
