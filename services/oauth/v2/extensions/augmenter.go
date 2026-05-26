package extensions

import (
	"encoding/json"
	"net/http"
)

// Augmenter is an extension point for adding the appropriate authorization information to oauth requests.
type Augmenter interface {
	// Augment adds the Authorization header to the request and sets the request body.
	Augment(r *http.Request, body []byte, secret json.RawMessage) error
}

type routerBodyAugmenter struct {
	AugmenterPath string
}
type routerHeaderAugmenter struct {
	HeaderName string
}
type headerAugmenter struct {
	HeaderName string
}

// RouterBodyAugmenter is an Augmenter that adds the authorization information to the request body.
var RouterBodyAugmenter = &routerBodyAugmenter{
	AugmenterPath: "input",
}

// HeaderAugmenter is an Augmenter that adds the authorization information to the request header.
var HeaderAugmenter = &headerAugmenter{
	HeaderName: "X-Rudder-Dest-Info",
}

var RouterHeaderAugmenter = &routerHeaderAugmenter{
	HeaderName: "Oauth-Secret",
}

// Augment adds the secret information to request body
func (t *routerBodyAugmenter) Augment(r *http.Request, body []byte, secret json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// Augment adds secret to request header to the request and sets the request body.
func (t *headerAugmenter) Augment(r *http.Request, body []byte, secret json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *routerHeaderAugmenter) Augment(r *http.Request, body []byte, secret json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}
