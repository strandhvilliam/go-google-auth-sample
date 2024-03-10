package infra

import "golang.org/x/oauth2"

type AuthProvider struct {
	*oauth2.Config
}

var authProvider = &AuthProvider{}

func GetAuthProvider() *AuthProvider {
	return authProvider
}

func NewAuthProvider(config func() *oauth2.Config) {
	res := config()
	authProvider.Config = res
}
