package config

import (
	"strconv"

	"github.com/zalando/go-keyring"
)

const (
	serviceName      = "stravastats"
	accessTokenName  = "accessTokenName"
	refreshTokenName = "refreshTokenName"
	expiresAtName    = "expiresAt"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int
}

func ReadTokens() (*Tokens, error) {
	accessToken, err := keyring.Get(serviceName, accessTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	refreshToken, err := keyring.Get(serviceName, refreshTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	expiresAt, err := keyring.Get(serviceName, expiresAtName)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	expiresAtInt, err := strconv.Atoi(expiresAt)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAtInt,
	}, nil
}

func SaveTokens(accessToken string, refreshToken string, expiresAt int) error {
	err := keyring.Set(serviceName, accessTokenName, accessToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, refreshTokenName, refreshToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, expiresAtName, strconv.Itoa(expiresAt))
	if err != nil {
		return err
	}

	return nil
}
