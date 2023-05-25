package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Rest interface {
	Get(url string, headers map[string]string, response any) (int, error)
}

type rest struct {
	client *http.Client
}

func NewRestClient(c *http.Client) Rest {
	return &rest{
		client: c,
	}
}

func (r *rest) Get(url string, headers map[string]string, response any) (int, error) {
	req, err := http.NewRequest(get, url, nil)
	if err != nil {
		return -1, fmt.Errorf("get: unable to create request: %s", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("get: unable to execute request: %s", err)
	}

	if res.Body != nil && response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return -1, fmt.Errorf("get: unable to decode response: %s", err)
		}
	}

	return res.StatusCode, nil
}