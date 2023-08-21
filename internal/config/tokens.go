package config

import (
	"stravastats/internal/model"
	"strconv"

	"github.com/zalando/go-keyring"
)

const (
	serviceName      = "stravastats"
	accessTokenName  = "accessTokenName"
	refreshTokenName = "refreshTokenName"
	expiresAtName    = "expiresAt"
)

func ReadTokens() (*model.Tokens, error) {
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

	return &model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAtInt,
	}, nil
}

func SaveTokens(tokens *model.Tokens) error {
	err := keyring.Set(serviceName, accessTokenName, tokens.AccessToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, refreshTokenName, tokens.RefreshToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, expiresAtName, strconv.Itoa(tokens.ExpiresAt))
	if err != nil {
		return err
	}

	return nil
}
