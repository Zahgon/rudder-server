package controlplane

import (
	"context"
	"errors"
	"net/http"
	"sync"
)

var ErrKeyNotFound = errors.New("request key not found")

type PublicPrivateKeyPair struct {
	PublicKey  string
	PrivateKey string
}

type BasicAuth struct {
	Username string
	Password string
}

type InternalControlPlane interface {
	GetDestinationSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error)
	GetSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error)
}

type internalClientWithCache struct {
	client InternalControlPlane
	cache  sync.Map
}

func NewInternalClient(baseURI string, auth BasicAuth) InternalControlPlane {
	_ = "STUB: not implemented"
	return *new(InternalControlPlane)
}

type internalClient struct {
	baseURI    string
	auth       BasicAuth
	httpClient *http.Client
}

func (api *internalClient) GetSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *internalClient) GetDestinationSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewInternalClientWithCache(baseURI string, auth BasicAuth) InternalControlPlane {
	_ = "STUB: not implemented"
	return *new(InternalControlPlane)
}

func (cc *internalClientWithCache) GetSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *internalClientWithCache) GetDestinationSSHKeys(ctx context.Context, id string) (*PublicPrivateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
