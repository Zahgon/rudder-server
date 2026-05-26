package augmenter

import (
	"encoding/json"
	"net/http"
)

type yandexAugmenter struct{}

var YandexReqAugmenter = &yandexAugmenter{}

// Custom augmenter for Yandex which sets token to Authorization header
func (y *yandexAugmenter) Augment(r *http.Request, body []byte, secret json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// format -> Authorization : OAuth <accessToken>

func GetAuthErrorCategoryForYandex(responseBody []byte) (string, error) {
	_ = "STUB: not implemented"
	/*
		Sample response for Yandex
		{
		    "errors": [
		        {
		            "error_type": "invalid_token",
		            "message": "Invalid oauth_token"
		        }
		    ],
		    "code": 403,
		    "message": "Invalid oauth_token"
		}
	*/return "", nil
}
