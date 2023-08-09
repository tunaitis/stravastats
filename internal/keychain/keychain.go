package keychain

import (
	"github.com/zalando/go-keyring"
)

const (
	serviceName = "stravastats"
	accessTokenName = "accessTokenName"
	refreshTokenName = "refreshTokenName"
)

func ReadTokens() (string, string, error) {
	accessToken, err := keyring.Get(serviceName, accessTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return "", "", err
	}

	refreshToken, err := keyring.Get(serviceName, refreshTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func WriteTokens(accessToken string, refreshToken string) error {
	err := keyring.Set(serviceName, accessTokenName, accessToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, refreshTokenName, refreshToken)
	return err
}

