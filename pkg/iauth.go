package pkg

import "net/http"

type IAuth interface {
	GetLwaTokenEndpoint() string

	GetDataEndpoint() string

	GetLwaCodeEndpoint() string

	GetLwaURL() string

	GetAccessTokenFromEndpoint() error

	ParseToken(resp *http.Response) error

	GetRefreshToken(code string) error

	BuildClient() *http.Client
}
