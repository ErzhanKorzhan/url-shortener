package api

import (
	"fmt"
	"net/http"
	"url-shortener/internal/storage"
)

func GetRedirect(url string) (string, error) {
	const op = "api.GetRedirect"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusFound {
		return "", fmt.Errorf("%s: %w: %d", op, storage.ErrInvalidStatusCode, resp.StatusCode)
	}

	return resp.Header.Get("Location"), nil
}

func DeleteStatusCode(url string) (int, error) {
	const op = "api.DeleteRedirect"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer func() { _ = resp.Body.Close() }()

	return resp.StatusCode, nil
}
