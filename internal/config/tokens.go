package config

import (
	"stravastats/internal/api"
	"strconv"

	"github.com/zalando/go-keyring"
)

const (
	serviceName      = "stravastats"
	accessTokenName  = "accessTokenName"
	refreshTokenName = "refreshTokenName"
	expiresAt        = "expiresAt"
)

func ReadTokens() (*api.TokensResponse, error) {
	accessToken, err := keyring.Get(serviceName, accessTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	refreshToken, err := keyring.Get(serviceName, refreshTokenName)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	expiresAt, err := keyring.Get(serviceName, expiresAt)
	if err != nil && err != keyring.ErrNotFound {
		return nil, err
	}

	expiresAtInt, err := strconv.Atoi(expiresAt)
	if err != nil {
		return nil, err
	}

	return &api.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAtInt,
	}, nil
}

func SaveTokens(tokens *api.TokensResponse) error {
	err := keyring.Set(serviceName, accessTokenName, tokens.AccessToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, refreshTokenName, tokens.RefreshToken)
	if err != nil {
		return err
	}

	err = keyring.Set(serviceName, expiresAt, strconv.Itoa(tokens.ExpiresAt))
	if err != nil {
		return err
	}

	return nil
}
